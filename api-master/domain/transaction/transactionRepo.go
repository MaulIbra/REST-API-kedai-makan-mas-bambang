package transaction

import "github.com/maulIbra/clean-architecture-go/api-master/domain/menu"

type ITransactionRepo interface {
	PostTransaction(transaction *Transaction,x map[string]int) error
	GetTransaction(counter string) ([]*TransactionResponseTemp, error)
	GetTransactionByID(id string) ([]TransactionResponseTemp, error)
	CheckMenuStock(id string) (*menu.Menu,error)
}
