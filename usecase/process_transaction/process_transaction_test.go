package process_transaction

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	mock_broker "github.com/herculesgabriel/payments-processor/adapter/broker/mock"
	"github.com/herculesgabriel/payments-processor/domain/entity"
	mock_repository "github.com/herculesgabriel/payments-processor/domain/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestProcessTransaction_ExecuteInvalidCreditCard(t *testing.T) {
	input := TransactionDTOInput{
		ID:                        "1",
		AccountID:                 "1",
		Amount:                    100,
		CreditCardNumber:          "1234567891324567",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
	}

	expectedOutput := TransactionDTOOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "invalid credit card number",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	producerMock := mock_broker.NewMockProducerInterface(ctrl)
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.ID), "transaction_processed")

	usecase := NewProcessTransaction(repositoryMock, producerMock, "transaction_processed")
	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestProcessTransaction_ExecuteRejectedTransaction(t *testing.T) {
	input := TransactionDTOInput{
		ID:                        "1",
		AccountID:                 "1",
		Amount:                    1200,
		CreditCardNumber:          "4796110246244203",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
	}

	expectedOutput := TransactionDTOOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "you do not have enough limit for this transaction",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	producerMock := mock_broker.NewMockProducerInterface(ctrl)
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.ID), "transaction_processed")

	usecase := NewProcessTransaction(repositoryMock, producerMock, "transaction_processed")
	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestProcessTransaction_ExecuteApprovedTransaction(t *testing.T) {
	input := TransactionDTOInput{
		ID:                        "1",
		AccountID:                 "1",
		Amount:                    980,
		CreditCardNumber:          "4796110246244203",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
	}

	expectedOutput := TransactionDTOOutput{
		ID:           "1",
		Status:       entity.APPROVED,
		ErrorMessage: "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	producerMock := mock_broker.NewMockProducerInterface(ctrl)
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.ID), "transaction_processed")

	usecase := NewProcessTransaction(repositoryMock, producerMock, "transaction_processed")
	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}
