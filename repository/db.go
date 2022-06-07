package repository

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(opts *DBOptions) (*gorm.DB, error) {
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
