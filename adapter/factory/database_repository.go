package factory

import (
	"database/sql"

	"github.com/herculesgabriel/payments-processor/adapter/repository"
)

type RepositoryDatabaseFactory struct {
	DB *sql.DB
}

func NewRepositoryDatabaseFactory(db *sql.DB) *RepositoryDatabaseFactory {
	return &RepositoryDatabaseFactory{DB: db}
}

func (r *RepositoryDatabaseFactory) CreateTransactionRepository() repository.TransactionRepositoryDB {
	return *repository.NewTransactionRepositoryDB(r.DB)
}
