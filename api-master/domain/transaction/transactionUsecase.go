package transaction

type ITransactionUsecase interface {
	PostTransaction(transaction *Transaction) (*string,error)
	GetTransactionByID(id string) (*TransactionResponse, error)
	GetTransaction(counter string) ([]*TransactionResponse, error)
}
