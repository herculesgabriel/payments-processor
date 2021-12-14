package entity

import "errors"

const (
	REJECTED = "rejected"
	APPROVED = "approved"
)

type Transaction struct {
	ID           string
	AccountID    string
	Amount       float64
	CreditCard   CreditCard
	Status       string
	ErrorMessage string
}

func (t *Transaction) IsValid() error {
	if t.Amount > 1000 {
		return errors.New("you do not have enough limit for this transaction")
	}
	if t.Amount < 1 {
		return errors.New("the amount must be greater than 1")
	}
	return nil
}

func (t *Transaction) SetCreditCard(card CreditCard) {
	t.CreditCard = card
}

func NewTransaction() *Transaction {
	return &Transaction{}
}
