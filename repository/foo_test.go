package repository_test

import (
	"testing"

	"github.com/emiliomarin/go-repository/repository"
	"github.com/stretchr/testify/assert"
)

// Scenario: I want to
// - Create a Foo object
// - Update its value
// - Set the flag as active
// If any fails rollback

func TestRegularFlow(t *testing.T) {
	// Without Txs is not possible to rollback
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

func TestTXFlow(t *testing.T) {
	db, err := NewTestDB()
	if err != nil {
		t.FailNow()
	}

	repo := repository.NewFooRepo(db)

	tx := repo.DB().Begin()
	initialValue := "initial-value"
	foo, err := repo.WithTx(tx).CreateFoo(initialValue)
	if err != nil {
		assert.Nil(t, err)
	}

	assert.Equal(t, foo.Value, initialValue, "should have the expected value")

	getFoo, err := repo.WithTx(tx).GetFoo(foo.ID)
	if err != nil {
		assert.Nil(t, err)
	}
	assert.Equal(t, getFoo.ID, foo.ID, "should have the same ID")
	assert.Equal(t, getFoo.Value, foo.Value, "should have the same value")

	updatedValue := "updated-value"
	updatedFoo, err := repo.WithTx(tx).UpdateValue(foo.ID, updatedValue)
	if err != nil {
		assert.Nil(t, err)
	}
	assert.Equal(t, updatedFoo.ID, foo.ID, "should have the same ID")
	assert.Equal(t, updatedFoo.Value, updatedValue, "should have updated the value")

	err = tx.Commit().Error
	assert.Nil(t, err, "should commit successfully")

	// Can we use the commited DB instance that is set on the struct?
	// No, we can't re-use a committed or rolled tx
	newFoo, err := repo.CreateFoo("new-foo")
	assert.Equal(t, err.Error(), "sql: transaction has already been committed or rolled back")
	assert.Nil(t, newFoo)
}
