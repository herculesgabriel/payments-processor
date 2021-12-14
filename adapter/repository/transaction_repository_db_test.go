package repository

import (
	"os"
	"testing"

	"github.com/herculesgabriel/payments-processor/adapter/repository/fixture"
	"github.com/herculesgabriel/payments-processor/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestTransactionRepositoryDBInsert(t *testing.T) {
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)

	repository := NewTransactionRepositoryDB(db)
	err := repository.Insert("1", "1", 12.1, entity.APPROVED, "")

	assert.Nil(t, err)
}
