package transaction

import (
	"encoding/json"

	"github.com/herculesgabriel/payments-processor/usecase/process_transaction"
)

type KafkaTransactionPresenter struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

func NewKafkaTransactionPresenter() *KafkaTransactionPresenter {
	return &KafkaTransactionPresenter{}
}

func (k *KafkaTransactionPresenter) Bind(input interface{}) error {
	k.ID = input.(process_transaction.TransactionDTOOutput).ID
	k.Status = input.(process_transaction.TransactionDTOOutput).Status
	k.ErrorMessage = input.(process_transaction.TransactionDTOOutput).ErrorMessage

	return nil
}

func (k *KafkaTransactionPresenter) Show() ([]byte, error) {
	j, err := json.Marshal(k)

	if err != nil {
		return nil, err
	}

	return j, nil
}
