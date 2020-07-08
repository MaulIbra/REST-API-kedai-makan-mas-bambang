package transaction

import "github.com/maulIbra/clean-architecture-go/api-master/models"

type ITransactionUsecase interface {
	PostTransaction(transaction *models.Transaction) (*string,error)
	GetTransactionByID(id string) (*models.TransactionResponse, error)
	GetTransaction(counter string) ([]*models.TransactionResponse, error)
}
