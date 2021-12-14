package repository

import (
	"database/sql"
	"time"
)

type TransactionRepositoryDB struct {
	db *sql.DB
}

func NewTransactionRepositoryDB(db *sql.DB) *TransactionRepositoryDB {
	return &TransactionRepositoryDB{db: db}
}

func (t *TransactionRepositoryDB) Insert(id string, account string, amount float64, status string, errorMessage string) error {
	stmt, err := t.db.Prepare(`
		INSERT INTO transactions (id, account_id, amount, status, error_message, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
	`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(id, account, amount, status, errorMessage, time.Now(), time.Now())

	if err != nil {
		return err
	}

	return nil
}
