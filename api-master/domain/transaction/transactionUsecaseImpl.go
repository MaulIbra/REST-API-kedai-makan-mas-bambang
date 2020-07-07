package transaction

import (
	"fmt"
	"github.com/maulIbra/clean-architecture-go/utils"
)

type TransactionUsecase struct {
	repo ITransactionRepo
}

func NewTransactionUsecase(repo ITransactionRepo) ITransactionUsecase{
	return &TransactionUsecase{
		repo: repo,
	}
}

func (t TransactionUsecase) GetTransaction(counter string) ([]*TransactionResponse, error) {
	transaction := []*TransactionResponse{}
	transactionTemp,err := t.repo.GetTransaction(counter)
	if err != nil {
		return nil, err
	}
	if len(transactionTemp) > 0 {
		transactionId := transactionTemp[0].TransactionId
		temp := []TransactionResponseTemp{}
		for id,val := range transactionTemp{
			if val.TransactionId == transactionId{
				temp = append(temp,*val)
			}else{
				listMenu,totalPrice,date := BundleListMenuTransaction(temp)
				transaction = append(transaction,&TransactionResponse{
					TransactionId:transactionId,
					Listmenu: listMenu,
					TransactionDate: date,
					TransactionTotalPrice:totalPrice,
				})
				temp = []TransactionResponseTemp{}
				temp = append(temp,*val)
				transactionId = val.TransactionId
			}
			if id == len(transactionTemp)-1{
				listMenu,totalPrice,date := BundleListMenuTransaction(temp)
				transaction = append(transaction,&TransactionResponse{
					TransactionId:transactionId,
					Listmenu: listMenu,
					TransactionDate: date,
					TransactionTotalPrice:totalPrice,
				})
			}
		}
	}
	return transaction,nil
}



func (t TransactionUsecase) PostTransaction(transaction *Transaction) (*string,error) {
	transaction.TransactionDate = utils.GetTimeNow()
	updateStock := make(map[string]int)
	for _,val := range transaction.ListMenu{
		menu,err := t.repo.CheckMenuStock(val.MenuId)
		if err != nil {
			return nil,err
		}
		if menu.Stock < val.Quantity{
			message := fmt.Sprintf("%v Sementara Kosong",menu.MenuName)
			return &message,nil
		}
		updateStock[val.MenuId]= menu.Stock-val.Quantity
	}
	if len(updateStock) > 0 {
		err := t.repo.PostTransaction(transaction,updateStock)
		if err != nil {
			return nil,err
		}
	}

	return nil,nil
}

func (t TransactionUsecase) GetTransactionByID(id string) (*TransactionResponse, error) {
	transactionTemp , err := t.repo.GetTransactionByID(id)
	if err != nil || len(transactionTemp)==0 {
		return nil, err
	}
	listMenu,totalPrice,_ := BundleListMenuTransaction(transactionTemp)
	transaction := TransactionResponse{
		TransactionId:transactionTemp[0].TransactionId,
		Listmenu: listMenu,
		TransactionDate: transactionTemp[0].TransactionDate,
		TransactionTotalPrice:totalPrice,
	}
	return &transaction , nil
}

func BundleListMenuTransaction(transactionTemp []TransactionResponseTemp) ([]*TransactionMenuResponse, int,string) {
	transactionMenu := []*TransactionMenuResponse{}
	var transactionTotalPrice int
	var transactionDate string
	for _,val := range transactionTemp {
		transactionMenu = append(transactionMenu,&TransactionMenuResponse{
			MenuId: val.Menu.MenuId,
			MenuName: val.Menu.MenuName,
			Quantity: val.Menu.Quantity,
			MenuPrice: val.Menu.MenuPrice,
			TotalPrice: val.Menu.TotalPrice,
		})
		transactionTotalPrice = transactionTotalPrice + val.Menu.TotalPrice
		transactionDate = val.TransactionDate
	}
	return transactionMenu,transactionTotalPrice,transactionDate
}
