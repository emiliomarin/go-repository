package repository

import (
	"errors"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

const barTable = "bar"

type Bar struct {
	ID    uuid.UUID
	Value string
}

type barRepo struct {
	db *gorm.DB
}

func NewBarRepo(db *gorm.DB) *barRepo {
	return &barRepo{
		db: db,
	}
}

type IBar interface {
	GetBar(id uuid.UUID) (*Bar, error)
	CreateBar(value string) (*Bar, error)
	UpdateValue(id uuid.UUID, value string) (*Bar, error)
}

func (r *barRepo) GetBar(id uuid.UUID) (*Bar, error) {
	var res *Bar
	if err := r.db.Table(barTable).Where("id = ?", id.String()).First(&res).Error; err != nil {
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

	err := r.db.Table(barTable).Create(f).Error
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (r *barRepo) UpdateValue(id uuid.UUID, value string) (*Bar, error) {
	f := &Bar{
		ID:    id,
		Value: value,
	}

	err := r.db.Table(barTable).Save(f).Error
	if err != nil {
		return nil, err
	}
	return f, nil
}
