package transaction

type ITransactionRepo interface {
	PostTransaction(transaction *Transaction) error
	GetTransaction(counter string) ([]*TransactionResponseTemp, error)
	GetTransactionByID(id string) ([]TransactionResponseTemp, error)
}
