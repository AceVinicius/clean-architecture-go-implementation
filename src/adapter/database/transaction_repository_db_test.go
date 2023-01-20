package database

import (
	"clean_architecture/src/adapter/database/fixture"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionRepositoryDatabaseInsert(t *testing.T) {
	migrationsDir := os.DirFS("fixture/sql")

	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)

	repository := NewTransactionRepositoryDb(db)
	err := repository.Insert("1", "1", 2, "approved", "")

	assert.Nil(t, err)
}
