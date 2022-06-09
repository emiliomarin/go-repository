package repository

import "gorm.io/gorm"

type IFooGen interface {
	ICRUD[Foo]
	// What if we want some custom functions
	// beside the basic crud? It wouldn't work
}

func NewFooRepoGen(db *gorm.DB) IFooGen {
	return NewCRUD[Foo](db)
}
