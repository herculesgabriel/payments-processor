package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("1234567891324567", "Hercules Gabriel", 11, 2029, 123)
	assert.Equal(t, "invalid credit card number", err.Error())

	_, err = NewCreditCard("4796110246244203", "Hercules Gabriel", 11, 2029, 123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationMonth(t *testing.T) {
	_, err := NewCreditCard("4796110246244203", "Hercules Gabriel", 13, 2029, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("4796110246244203", "Hercules Gabriel", 0, 2029, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("4796110246244203", "Hercules Gabriel", 12, 2029, 123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationYear(t *testing.T) {
	lastYearDate := time.Now().AddDate(-1, 0, 0)
	_, err := NewCreditCard("4796110246244203", "Hercules Gabriel", 12, lastYearDate.Year(), 123)
	assert.Equal(t, "invalid expiration year", err.Error())

	currentYearDate := time.Now()
	_, err = NewCreditCard("4796110246244203", "Hercules Gabriel", 12, currentYearDate.Year(), 123)
	assert.Nil(t, err)

	nextYearDate := time.Now().AddDate(1, 0, 0)
	_, err = NewCreditCard("4796110246244203", "Hercules Gabriel", 12, nextYearDate.Year(), 123)
	assert.Nil(t, err)
}
