package repository

import (
	"context"
	"database/sql"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) Create(ctx context.Context, id int64, balance string) error {
	query := `
		INSERT INTO accounts (id, balance)
		VALUES ($1, $2)
	`

	_, err := r.db.ExecContext(ctx, query, id, balance)
	return err
}

func (r *AccountRepository) GetByID(ctx context.Context, id int64) (int64, string, error) {
	query := `
		SELECT id, balance
		FROM accounts
		WHERE id = $1
	`

	var accountID int64
	var balance string

	err := r.db.QueryRowContext(ctx, query, id).Scan(&accountID, &balance)
	if err != nil {
		return 0, "", err
	}

	return accountID, balance, nil
}

