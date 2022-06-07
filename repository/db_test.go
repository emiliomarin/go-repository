package repository_test

import (
	"github.com/emiliomarin/go-repository/repository"
	"gorm.io/gorm"
)

// NewTestDB will be used to create a test DB connection
// and teardown logic for cleanup
func NewTestDB() (*gorm.DB, error) {
	opts := &repository.DBOptions{
		Host:     "127.0.0.1",
		Port:     "5434",
		User:     "arexdb_dev",
		Password: "arexdb_dev",
		Database: "test_db",
	}

	return repository.NewDB(opts)
}
