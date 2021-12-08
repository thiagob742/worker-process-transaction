package entity

import (

	"errors"
)

type Transaction struct {

	ID string
	amout float64
	accountID string
	Status string
	ErrorMessage string
	CreditCard CreditCard

}

func NewTransaction()*Transaction {
	return &Transaction{}
}

func (t *Transaction) IsValid() error{

	err := t.AmountIsNotMax()

	if err != nil{
		return err
	}
	err = t.AmountIsNotMin()

	if err != nil{
		return err
	}

	return nil
}

func (t *Transaction) AmountIsNotMax() error {

	if (t.amout < 1000){
		
		return	nil
	}

	return errors.New("value max of transaction")

}

func (t *Transaction) AmountIsNotMin() error {

	if (t.amout >= 1){
		
		return	nil
	}

	return errors.New("value min of transaction")

}

func (t *Transaction) SetCreditCard(card CreditCard) {
	t.CreditCard = card
}