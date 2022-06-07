package repository

import (
	"errors"
	"log"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

const barTable = "bar"

type Bar struct {
	ID    uuid.UUID
	Value string
	Flag  bool
}

type barRepo struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewBarRepo(db *gorm.DB) *barRepo {
	return &barRepo{
		db: db,
	}
}

func (r *barRepo) GetBar(id uuid.UUID) (*Bar, error) {
	var res *Bar
	if err := r.DB().Table(barTable).Where("id = ?", id.String()).First(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return res, nil
}

func (r *barRepo) CreateBar(value string) (*Bar, error) {
	f := &Bar{
		ID:    uuid.Must(uuid.NewV4()),
		Value: value,
	}

	err := r.DB().Table(barTable).Create(f).Error
	if err != nil {
		log.Println("Error creating", err)
		return nil, err
	}
	return f, nil
}

func (r *barRepo) UpdateValue(id uuid.UUID, value string) (*Bar, error) {
	f := &Bar{
		ID:    id,
		Value: value,
	}

	err := r.DB().Table(barTable).Save(f).Error
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (r *barRepo) ActivateFlag(f *Bar) (*Bar, error) {
	f.Flag = true

	err := r.DB().Table(barTable).Save(f).Error
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (i *barRepo) WithTx(tx *gorm.DB) *barRepo {
	if tx == nil {
		return i
	}
	i.tx = tx
	return i
}

func (i *barRepo) DB() *gorm.DB {
	if i.tx != nil {
		return i.tx
	}
	return i.db
}

func (i *barRepo) Commit() error {
	err := i.tx.Commit().Error
	i.tx = nil
	return err
}
