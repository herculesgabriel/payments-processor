package entity

import (
	"errors"
	"regexp"
	"time"
)

type CreditCard struct {
	number          string
	name            string
	expirationMonth int
	expirationYear  int
	CVV             int
}

func (c *CreditCard) IsValid() error {
	err := c.ValidateNumber()
	if err != nil {
		return err
	}
	err = c.ValidateMonth()
	if err != nil {
		return err
	}
	err = c.ValidateYear()
	if err != nil {
		return err
	}

	return nil
}

func (c *CreditCard) ValidateNumber() error {
	validCardRegex := regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$`)
	if !validCardRegex.MatchString(c.number) {
		return errors.New("invalid credit card number")
	}
	return nil
}

func (c *CreditCard) ValidateMonth() error {
	if c.expirationMonth > 0 && c.expirationMonth < 13 {
		return nil
	}
	return errors.New("invalid expiration month")
}

func (c *CreditCard) ValidateYear() error {
	if c.expirationYear >= time.Now().Year() {
		return nil
	}
	return errors.New("invalid expiration year")
}

func NewCreditCard(number string, name string, expirationMonth int, expirationYear int, CVV int) (*CreditCard, error) {
	cc := &CreditCard{
		number:          number,
		name:            name,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		CVV:             CVV,
	}

	err := cc.IsValid()
	if err != nil {
		return nil, err
	}
	return cc, nil
}
