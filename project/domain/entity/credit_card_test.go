package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)


func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("40000000000000000", "Jose da Silva", 12, 2024, 123)
	assert.Equal(t, "invalid credit card number", err.Error())

	_, err = NewCreditCard("4193523830170205", "Jose da Silva", 12, 2024, 123)
	assert.Nil(t, err)
}


func TestCreditCardExpirationMonth(t *testing.T) {

	_, err := NewCreditCard("4193523830170205", "Jose da Silva", 12, 2024, 123)
	assert.Nil(t, err)

	_, err = NewCreditCard("4193523830170205", "Jose da Silva", 13, 2024, 123)
	assert.Equal(t, "invalid month, credit card has expired", err.Error())

	_, err = NewCreditCard("4193523830170205", "Jose da Silva", 0, 2024, 123)
	assert.Equal(t, "invalid month, credit card has expired", err.Error())

}

func TestCreditCardExpirationYear(t *testing.T){
	
	lastYear := time.Now().AddDate(-1, 0 , 0)

	_, err := NewCreditCard("4193523830170205", "Jose da Silva", 8, lastYear.Year(), 123)
	assert.Equal(t, "invalid year, credit card has expired", err.Error())

	_, err = NewCreditCard("4193523830170205", "Jose da Silva", 1, 2012, 123)
	assert.Equal(t, "invalid year, credit card has expired", err.Error())

	_, err = NewCreditCard("4193523830170205", "Jose da Silva", 12, time.Now().Year(), 123)
	assert.Nil(t, err)


}

