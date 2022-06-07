package repository_test

import (
	"errors"
	"testing"

	"github.com/emiliomarin/go-repository/repository"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUnitOfWork(t *testing.T) {
	db, err := NewTestDB()
	if err != nil {
		t.FailNow()
	}

	uow := repository.NewUnitOfWork(db)
	fooRepo := repository.NewFooRepo(db)
	barRepo := repository.NewBarRepo(db)

	initialValue := "initial-value"
	updatedValue := "updated-value"
	var fooID, barID uuid.UUID

	t.Run("Using a unit of work block on success", func(t *testing.T) {
		err = uow.Do(func(uows repository.UnitOfWorkStore) error {
			foo, err := uows.Foo().CreateFoo(initialValue)
			if err != nil {
				assert.Nil(t, err)
				return err
			}
			fooID = foo.ID
			getFoo, err := uows.Foo().GetFoo(foo.ID)
			if err != nil {
				assert.Nil(t, err)
				return err
			}
			assert.Equal(t, getFoo.ID, foo.ID, "should have the same ID")
			assert.Equal(t, getFoo.Value, foo.Value, "should have the same value")

			updatedFoo, err := uows.Foo().UpdateValue(foo.ID, updatedValue)
			if err != nil {
				assert.Nil(t, err)
				return err
			}
			assert.Equal(t, updatedFoo.ID, foo.ID, "should have the same ID")
			assert.Equal(t, updatedFoo.Value, updatedValue, "should have updated the value")

			return nil
		})
		assert.Nil(t, err)

		getFoo, err := fooRepo.GetFoo(fooID)
		if err != nil {
			assert.Nil(t, err)
		}
		assert.Equal(t, getFoo.ID, fooID)
		assert.Equal(t, getFoo.Value, updatedValue)

		t.Run("And we try to create another foo", func(t *testing.T) {
			f, err := fooRepo.CreateFoo(initialValue)
			assert.Nil(t, err)
			assert.Equal(t, f.Value, initialValue)
		})
	})

	t.Run("Rollback on error", func(t *testing.T) {
		err = uow.Do(func(uows repository.UnitOfWorkStore) error {
			foo, err := uows.Foo().CreateFoo(initialValue)
			if err != nil {
				assert.Nil(t, err)
				return err
			}
			fooID = foo.ID

			getFoo, err := uows.Foo().GetFoo(foo.ID)
			if err != nil {
				assert.Nil(t, err)
				return err
			}
			assert.Equal(t, getFoo.ID, foo.ID, "should have the same ID")
			assert.Equal(t, getFoo.Value, foo.Value, "should have the same value")

			updatedFoo, err := uows.Foo().UpdateValue(foo.ID, updatedValue)
			if err != nil {
				assert.Nil(t, err)
				return err
			}
			assert.Equal(t, updatedFoo.ID, foo.ID, "should have the same ID")
			assert.Equal(t, updatedFoo.Value, updatedValue, "should have updated the value")

			return errors.New("something went wrong")
		})
		assert.NotNil(t, err)

		getFoo, err := fooRepo.GetFoo(fooID)
		assert.Nil(t, err, "should not return an error")
		assert.Nil(t, getFoo, "should not have created foo")
	})

	t.Run("Multiple repos success", func(t *testing.T) {
		err = uow.Do(func(uows repository.UnitOfWorkStore) error {
			foo, err := uows.Foo().CreateFoo(initialValue)
			if err != nil {
				assert.Nil(t, err)
				return err
			}
			fooID = foo.ID

			bar, err := uows.Bar().CreateBar(initialValue)
			if err != nil {
				assert.Nil(t, err)
				return err
			}
			barID = bar.ID

			getFoo, err := uows.Foo().GetFoo(foo.ID)
			if err != nil {
				assert.Nil(t, err)
				return err
			}
			assert.Equal(t, getFoo.ID, foo.ID, "should have the same ID")
			assert.Equal(t, getFoo.Value, foo.Value, "should have the same value")

			getBar, err := uows.Bar().GetBar(bar.ID)
			if err != nil {
				assert.Nil(t, err)
				return err
			}
			assert.Equal(t, getBar.ID, bar.ID, "should have the same ID")
			assert.Equal(t, getBar.Value, bar.Value, "should have the same value")

			updatedFoo, err := uows.Foo().UpdateValue(foo.ID, updatedValue)
			if err != nil {
				assert.Nil(t, err)
				return err
			}
			assert.Equal(t, updatedFoo.ID, foo.ID, "should have the same ID")
			assert.Equal(t, updatedFoo.Value, updatedValue, "should have updated the value")

			updatedBar, err := uows.Bar().UpdateValue(bar.ID, updatedValue)
			if err != nil {
				assert.Nil(t, err)
				return err
			}
			assert.Equal(t, updatedBar.ID, bar.ID, "should have the same ID")
			assert.Equal(t, updatedBar.Value, updatedValue, "should have updated the value")

			return nil
		})
		assert.Nil(t, err)

		getFoo, err := fooRepo.GetFoo(fooID)
		if err != nil {
			assert.Nil(t, err)
		}
		assert.Equal(t, getFoo.ID, fooID)
		assert.Equal(t, getFoo.Value, updatedValue)

		getBar, err := barRepo.GetBar(barID)
		if err != nil {
			assert.Nil(t, err)
		}
		assert.Equal(t, getBar.ID, barID)
		assert.Equal(t, getBar.Value, updatedValue)

	})

	t.Run("Multiple repos rollback", func(t *testing.T) {
		err = uow.Do(func(uows repository.UnitOfWorkStore) error {
			foo, err := uows.Foo().CreateFoo(initialValue)
			if err != nil {
				assert.Nil(t, err)
				return err
			}
			fooID = foo.ID

			bar, err := uows.Bar().CreateBar(initialValue)
			if err != nil {
				assert.Nil(t, err)
				return err
			}
			barID = bar.ID

			getFoo, err := uows.Foo().GetFoo(foo.ID)
			if err != nil {
				assert.Nil(t, err)
				return err
			}
			assert.Equal(t, getFoo.ID, foo.ID, "should have the same ID")
			assert.Equal(t, getFoo.Value, foo.Value, "should have the same value")

			getBar, err := uows.Bar().GetBar(bar.ID)
			if err != nil {
				assert.Nil(t, err)
				return err
			}
			assert.Equal(t, getBar.ID, bar.ID, "should have the same ID")
			assert.Equal(t, getBar.Value, bar.Value, "should have the same value")

			return errors.New("something went wrong")
		})
		assert.NotNil(t, err, "should return an error")

		getFoo, err := fooRepo.GetFoo(fooID)
		assert.Nil(t, err)
		assert.Nil(t, getFoo, "should not have created foo")

		getBar, err := barRepo.GetBar(barID)
		assert.Nil(t, err)
		assert.Nil(t, getBar, "should not have created bar")
	})
}
