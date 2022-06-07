package repository_test

import (
	"testing"

	"github.com/emiliomarin/go-repository/repository"
	"github.com/stretchr/testify/assert"
)

func TestRegularFlow(t *testing.T) {
	db, err := NewTestDB()
	if err != nil {
		t.FailNow()
	}

	repo := repository.NewFooRepo(db)

	initialValue := "initial-value"
	foo, err := repo.CreateFoo(initialValue)
	if err != nil {
		assert.Nil(t, err)
	}

	assert.Equal(t, foo.Value, initialValue, "should have the expected value")

	getFoo, err := repo.GetFoo(foo.ID)
	if err != nil {
		assert.Nil(t, err)
	}
	assert.Equal(t, getFoo.ID, foo.ID, "should have the same ID")
	assert.Equal(t, getFoo.Value, foo.Value, "should have the same value")

	updatedValue := "updated-value"
	updatedFoo, err := repo.UpdateValue(foo.ID, updatedValue)
	if err != nil {
		assert.Nil(t, err)
	}
	assert.Equal(t, updatedFoo.ID, foo.ID, "should have the same ID")
	assert.Equal(t, updatedFoo.Value, updatedValue, "should have updated the value")
}
