package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionWithAmountGreaterThan1000(t *testing.T) {
	transaction := NewTransaction()

	transaction.Id = "1"
	transaction.AccountId = "1"
	transaction.Amount = 2000

	err := transaction.IsValid()

	assert.Error(t, err)
	assert.Equal(t, "you don't have limit for this transaction", err.Error())
}

func TestTransactionWithAmountEqualTo1000(t *testing.T) {
	transaction := NewTransaction()

	transaction.Id = "1"
	transaction.AccountId = "1"
	transaction.Amount = 1000

	err := transaction.IsValid()

	assert.Nil(t, err)
	assert.Equal(t, nil, err)
}

func TestTransactionWithAmountBetween1And1000(t *testing.T) {
	transaction := NewTransaction()

	transaction.Id = "1"
	transaction.AccountId = "1"
	transaction.Amount = 350

	err := transaction.IsValid()

	assert.Nil(t, err)
	assert.Equal(t, nil, err)
}

func TestTransactionWithAmountEqualTo1(t *testing.T) {
	transaction := NewTransaction()

	transaction.Id = "1"
	transaction.AccountId = "1"
	transaction.Amount = 1

	err := transaction.IsValid()

	assert.Nil(t, err)
	assert.Equal(t, nil, err)
}

func TestTransactionWithAmountSmallerThan1(t *testing.T) {
	transaction := NewTransaction()

	transaction.Id = "1"
	transaction.AccountId = "1"
	transaction.Amount = 0.3

	err := transaction.IsValid()

	assert.Error(t, err)
	assert.Equal(t, "the amount must be greater than 1", err.Error())
}

func TestTransactionWithAmountNegative(t *testing.T) {
	transaction := NewTransaction()

	transaction.Id = "1"
	transaction.AccountId = "1"
	transaction.Amount = -8

	err := transaction.IsValid()

	assert.Error(t, err)
	assert.Equal(t, "the amount must be greater than 1", err.Error())
}
