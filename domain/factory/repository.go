package factory

import "github.com/herculesgabriel/payments-processor/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
