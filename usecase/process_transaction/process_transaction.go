package process_transaction

import (
	"github.com/herculesgabriel/payments-processor/adapter/broker"
	"github.com/herculesgabriel/payments-processor/domain/entity"
	"github.com/herculesgabriel/payments-processor/domain/repository"
)

type ProcessTransaction struct {
	Repository repository.TransactionRepository
	Producer   broker.ProducerInterface
	Topic      string
}

func NewProcessTransaction(repository repository.TransactionRepository, producerInterface broker.ProducerInterface, topic string) *ProcessTransaction {
	return &ProcessTransaction{Repository: repository, Producer: producerInterface, Topic: topic}
}

func (p *ProcessTransaction) RejectTransaction(transaction *entity.Transaction, err error) (TransactionDTOOutput, error) {
	insertError := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, entity.REJECTED, err.Error())

	if insertError != nil {
		return TransactionDTOOutput{}, insertError
	}

	output := TransactionDTOOutput{
		ID:           transaction.ID,
		Status:       entity.REJECTED,
		ErrorMessage: err.Error(),
	}

	publishError := p.publish(output, []byte(output.ID))

	if publishError != nil {
		return TransactionDTOOutput{}, publishError
	}

	return output, nil
}

func (p *ProcessTransaction) ApproveTransaction(transaction *entity.Transaction) (TransactionDTOOutput, error) {
	insertError := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, entity.APPROVED, "")

	if insertError != nil {
		return TransactionDTOOutput{}, insertError
	}

	output := TransactionDTOOutput{
		ID:           transaction.ID,
		Status:       entity.APPROVED,
		ErrorMessage: "",
	}

	publishError := p.publish(output, []byte(output.ID))

	if publishError != nil {
		return TransactionDTOOutput{}, publishError
	}

	return output, nil
}

func (p *ProcessTransaction) Execute(input TransactionDTOInput) (TransactionDTOOutput, error) {
	transaction := entity.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount
	creditCard, invalidCreditCard := entity.NewCreditCard(input.CreditCardNumber, input.CreditCardName, input.CreditCardExpirationMonth, input.CreditCardExpirationYear, input.CreditCardCVV)

	if invalidCreditCard != nil {
		return p.RejectTransaction(transaction, invalidCreditCard)
	}

	transaction.SetCreditCard(*creditCard)
	invalidTransaction := transaction.IsValid()

	if invalidTransaction != nil {
		return p.RejectTransaction(transaction, invalidTransaction)
	}

	return p.ApproveTransaction(transaction)
}

func (p *ProcessTransaction) publish(output TransactionDTOOutput, key []byte) error {
	err := p.Producer.Publish(output, key, p.Topic)

	if err != nil {
		return err
	}

	return nil
}
