package repository

import (
	"errors"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type ICRUD[T any] interface {
	Get(id uuid.UUID) (*T, error)
	Create(model *T) error
	Update(id uuid.UUID, value string) (*T, error)
}

type CRUD[T any] struct {
	db *gorm.DB
}

func NewCRUD[T any](db *gorm.DB) ICRUD[T] {
	return CRUD[T]{db: db}
}

func (c CRUD[T]) Get(id uuid.UUID) (*T, error) {
	var res *T
	if err := c.db.Model(&res).Where("id = ?", id.String()).First(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return res, nil

}

// Como podemos cambiar la tabla?
// Podemos meter la funcion TableName pero acoplamos capas
func (c CRUD[T]) Create(model *T) error {
	return nil
}

func (c CRUD[T]) Update(id uuid.UUID, value string) (*T, error) {
	return nil, nil
}
