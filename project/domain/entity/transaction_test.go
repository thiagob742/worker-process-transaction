package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/google/uuid"
)


func TestTransactionIsValid(t *testing.T){

	transaction := NewTransaction()
	transaction.ID = uuid.New().String()
	transaction.amout = 900
	transaction.accountID = uuid.New().String()

	assert.Nil(t, transaction.IsValid())

}

func TestTransactionValueMax(t *testing.T) {

	transaction:= NewTransaction()

	transaction.ID = uuid.New().String()
	transaction.amout = 1001
	transaction.accountID = uuid.New().String()


	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "value max of transaction", err.Error())

}

func TestTransactionValueMin(t *testing.T) {

	transaction:= NewTransaction()

	transaction.ID = uuid.New().String()
	transaction.amout = -1
	transaction.accountID = uuid.New().String()


	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "value min of transaction", err.Error())

}