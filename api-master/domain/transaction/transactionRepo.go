package transaction

import (
	"github.com/maulIbra/clean-architecture-go/api-master/models"
)

type ITransactionRepo interface {
	PostTransaction(transaction *models.Transaction,x map[string]int) error
	GetTransaction(counter string) ([]*models.TransactionResponseTemp, error)
	GetTransactionByID(id string) ([]models.TransactionResponseTemp, error)
	CheckMenuStock(id string) (*models.Menu,error)
}
