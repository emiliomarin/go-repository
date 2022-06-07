package repository_test

import (
	"testing"

	"github.com/emiliomarin/go-repository/repository"
	"github.com/stretchr/testify/assert"
)

func TestBar(t *testing.T) {
	db, err := NewTestDB()
	if err != nil {
		t.FailNow()
	}

	repo := repository.NewBarRepo(db)

	initialValue := "initial-value"
	foo, err := repo.CreateBar(initialValue)
	if err != nil {
		assert.Nil(t, err)
	}

	assert.Equal(t, foo.Value, initialValue, "should have the expected value")

	getBar, err := repo.GetBar(foo.ID)
	if err != nil {
		assert.Nil(t, err)
	}
	assert.Equal(t, getBar.ID, foo.ID, "should have the same ID")
	assert.Equal(t, getBar.Value, foo.Value, "should have the same value")

	updatedValue := "updated-value"
	updatedBar, err := repo.UpdateValue(foo.ID, updatedValue)
	if err != nil {
		assert.Nil(t, err)
	}
	assert.Equal(t, updatedBar.ID, foo.ID, "should have the same ID")
	assert.Equal(t, updatedBar.Value, updatedValue, "should have updated the value")
}
