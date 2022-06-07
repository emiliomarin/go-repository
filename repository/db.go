package repository

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(opts *DBOptions) (*gorm.DB, error) {
	// opts := &DBOptions{
	// 	Host:     "127.0.0.1",
	// 	Port:     "5432",
	// 	User:     "arexdb_dev",
	// 	Password: "arexdb_dev",
	// 	Database: "test_db",
	// 	SSLMode:  "disable",
	// }

	db, err := gorm.Open(postgres.Open(opts.connection()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// DBOptions defines the data needed to create a db connection
// SSLMode is a optional option that will take disable as default
type DBOptions struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (o *DBOptions) connection() string {
	return fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		o.Host,
		o.Port,
		o.Database,
		o.User,
		o.Password,
		o.SSLMode,
	)
}

// type Tx struct {
// 	tx *gorm.DB
// }

// func NewTx(db *gorm.DB) Tx {
// 	return Tx{tx: db}
// }

// func (t Tx) Begin() *gorm.DB {
// 	return t.tx.Begin()
// }

// func (t Tx) Commit() error {
// 	log.Println(t.tx)
// 	return t.tx.Commit().Error
// }

// func (t Tx) Rollback() error {
// 	return t.tx.Rollback().Error
// }
