package repository

import (
	"context"
	"database/sql"
	"github.com/wellingtonbrunodev/internal_transfer_system_with_golang/internal/domain"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Transfer(
	ctx context.Context,
	sourceID int64,
	destinationID int64,
	amount string,
) error {

	if sourceID == destinationID {
		return domain.ErrSameAccountTransfer
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Lock accounts in deterministic order to avoid deadlocks
	firstID := sourceID
	secondID := destinationID

	if sourceID > destinationID {
		firstID = destinationID
		secondID = sourceID
	}

	lockQuery := `
		SELECT id
		FROM accounts
		WHERE id = $1 OR id = $2
		FOR UPDATE
	`

	rows, err := tx.QueryContext(ctx, lockQuery, firstID, secondID)
	if err != nil {
		return err
	}
	defer rows.Close()

	found := make(map[int64]bool)
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return err
		}
		found[id] = true
	}

	if !found[sourceID] || !found[destinationID] {
		return domain.ErrAccountNotFound
	}

	// Check source balance
	var sourceBalance string
	err = tx.QueryRowContext(
		ctx,
		`SELECT balance FROM accounts WHERE id = $1`,
		sourceID,
	).Scan(&sourceBalance)
	if err != nil {
		return err
	}

	// Compare balances numerically inside DB
	var sufficient bool
	err = tx.QueryRowContext(
		ctx,
		`SELECT $1::numeric <= balance FROM accounts WHERE id = $2`,
		amount,
		sourceID,
	).Scan(&sufficient)
	if err != nil {
		return err
	}

	if !sufficient {
		return domain.ErrInsufficientBalance
	}

	// Deduct from source
	_, err = tx.ExecContext(
		ctx,
		`UPDATE accounts SET balance = balance - $1 WHERE id = $2`,
		amount,
		sourceID,
	)
	if err != nil {
		return err
	}

	// Add to destination
	_, err = tx.ExecContext(
		ctx,
		`UPDATE accounts SET balance = balance + $1 WHERE id = $2`,
		amount,
		destinationID,
	)
	if err != nil {
		return err
	}

	// Insert transaction log
	_, err = tx.ExecContext(
		ctx,
		`INSERT INTO transactions (type, source_account_id, destination_account_id, amount)
		 VALUES ('TRANSFER', $1, $2, $3)`,
		sourceID,
		destinationID,
		amount,
	)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
