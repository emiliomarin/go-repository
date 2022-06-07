package repository

import (
	"gorm.io/gorm"
)

type uowStore struct {
	foo IFoo
	bar IBar
}

// UnitOfWorkStore provides access to datastores that can be
// used inside an Unit-of-Work. All data changes done through
// them will be executed inside a DB transaction
type UnitOfWorkStore interface {
	Foo() IFoo
	Bar() IBar
}

func (u uowStore) Foo() IFoo {
	return u.foo
}

func (u uowStore) Bar() IBar {
	return u.bar
}

type UnitOfWorkBlock func(UnitOfWorkStore) error

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
		newStore := &uowStore{
			foo: NewFooRepo(tx),
			bar: NewBarRepo(tx),
		}
		return fn(newStore)
	})
}
