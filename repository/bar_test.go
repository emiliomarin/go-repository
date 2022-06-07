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

	tx := repo.DB().Begin()
	initialValue := "initial-value"
	bar, err := repo.WithTx(tx).CreateBar(initialValue)
	if err != nil {
		assert.Nil(t, err)
	}

	assert.Equal(t, bar.Value, initialValue, "should have the expected value")

	getBar, err := repo.WithTx(tx).GetBar(bar.ID)
	if err != nil {
		assert.Nil(t, err)
	}
	assert.Equal(t, getBar.ID, bar.ID, "should have the same ID")
	assert.Equal(t, getBar.Value, bar.Value, "should have the same value")

	updatedValue := "updated-value"
	updatedBar, err := repo.WithTx(tx).UpdateValue(bar.ID, updatedValue)
	if err != nil {
		assert.Nil(t, err)
	}
	assert.Equal(t, updatedBar.ID, bar.ID, "should have the same ID")
	assert.Equal(t, updatedBar.Value, updatedValue, "should have updated the value")

	err = repo.Commit()
	assert.Nil(t, err, "should commit successfully")

	// Can we use the commited DB instance that is set on the struct?
	// Now that the repo tx is set to nil yes. But it wouldn't work with 2 repos
	newBar, err := repo.CreateBar("new-bar")
	if err != nil {
		assert.Nil(t, err)
	}
	assert.Equal(t, newBar.Value, "new-bar")
}
