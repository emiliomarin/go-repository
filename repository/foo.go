package repository

import (
	"errors"
	"log"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

const fooTable = "foo"

type Foo struct {
	ID    uuid.UUID
	Value string
	Flag  bool
}

type fooRepo struct {
	db *gorm.DB
}

func NewFooRepo(db *gorm.DB) *fooRepo {
	return &fooRepo{
		db: db,
	}
}

func (r *fooRepo) GetFoo(id uuid.UUID) (*Foo, error) {
	var res *Foo
	if err := r.db.Table(fooTable).Where("id = ?", id.String()).First(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return res, nil
}

func (r *fooRepo) CreateFoo(value string) (*Foo, error) {
	f := &Foo{
		ID:    uuid.Must(uuid.NewV4()),
		Value: value,
	}

	err := r.db.Table(fooTable).Create(f).Error
	if err != nil {
		log.Println("Error creating", err)
		return nil, err
	}
	return f, nil
}

func (r *fooRepo) UpdateValue(id uuid.UUID, value string) (*Foo, error) {
	f := &Foo{
		ID:    id,
		Value: value,
	}

	err := r.db.Table(fooTable).Save(f).Error
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (r *fooRepo) ActivateFlag(f *Foo) (*Foo, error) {
	f.Flag = true

	err := r.db.Table(fooTable).Save(f).Error
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (i *fooRepo) WithTx(tx *gorm.DB) *fooRepo {
	if tx == nil {
		return i
	}
	i.db = tx
	return i
}

func (i *fooRepo) DB() *gorm.DB {
	return i.db
}
