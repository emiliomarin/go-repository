package repository

import (
	"gorm.io/gorm"
)

type UnitOfWorkStore struct {
	Foo IFoo
	Bar IBar
}

type UnitOfWorkBlock func(*UnitOfWorkStore) error

type unitOfWork struct {
	db *gorm.DB
}

type UnitOfWork interface {
	Do(UnitOfWorkBlock) error
}

func NewUnitOfWork(db *gorm.DB) UnitOfWork {
	return &unitOfWork{db: db}
}

// Do executes the given UnitOfWorkBlock iniside a DB transaction
func (s *unitOfWork) Do(fn UnitOfWorkBlock) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		newStore := &UnitOfWorkStore{
			Foo: NewFooRepo(tx),
			Bar: NewBarRepo(tx),
		}
		return fn(newStore)
	})
}
