package repository

import (
	"context"
	"database/sql"
	"github.com/shopspring/decimal"
)

// AccountRepository implements AccountRepository interface.
type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

//Create inserts new account info to database
func (r *AccountRepository) Create(ctx context.Context, id int64, balance decimal.Decimal) error {
	query := `
		INSERT INTO accounts (id, balance)
		VALUES ($1, $2)
	`

	_, err := r.db.ExecContext(ctx, query, id, balance)
	return err
}

//GetByID reads account data from database by using the account id
func (r *AccountRepository) GetByID(ctx context.Context, id int64) (int64, decimal.Decimal, error) {
	query := `
		SELECT id, balance
		FROM accounts
		WHERE id = $1
	`

	var accountID int64
	var balanceStr string

	err := r.db.QueryRowContext(ctx, query, id).Scan(&accountID, &balanceStr)
	if err != nil {
		return 0, decimal.Zero, err
	}

	balance, err := decimal.NewFromString(balanceStr)
	if err != nil {
		return 0, decimal.Zero, err
	}

	return accountID, balance, nil
}

