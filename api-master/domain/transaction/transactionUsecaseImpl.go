package transaction

import (
	"fmt"
	"github.com/maulIbra/clean-architecture-go/api-master/models"
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

func (t TransactionUsecase) GetTransaction(counter string) ([]*models.TransactionResponse, error) {
	transaction := []*models.TransactionResponse{}
	transactionTemp,err := t.repo.GetTransaction(counter)
	if err != nil {
		return nil, err
	}
	if len(transactionTemp) > 0 {
		transactionId := transactionTemp[0].TransactionId
		temp := []models.TransactionResponseTemp{}
		for id,val := range transactionTemp{
			if val.TransactionId == transactionId{
				temp = append(temp,*val)
			}else{
				listMenu,totalPrice,date := BundleListMenuTransaction(temp)
				transaction = append(transaction,&models.TransactionResponse{
					TransactionId:transactionId,
					Listmenu: listMenu,
					TransactionDate: date,
					TransactionTotalPrice:totalPrice,
				})
				temp = []models.TransactionResponseTemp{}
				temp = append(temp,*val)
				transactionId = val.TransactionId
			}
			if id == len(transactionTemp)-1{
				listMenu,totalPrice,date := BundleListMenuTransaction(temp)
				transaction = append(transaction,&models.TransactionResponse{
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



func (t TransactionUsecase) PostTransaction(transaction *models.Transaction) (*string,error) {
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

func (t TransactionUsecase) GetTransactionByID(id string) (*models.TransactionResponse, error) {
	transactionTemp , err := t.repo.GetTransactionByID(id)
	if err != nil || len(transactionTemp)==0 {
		return nil, err
	}
	listMenu,totalPrice,_ := BundleListMenuTransaction(transactionTemp)
	transaction := models.TransactionResponse{
		TransactionId:transactionTemp[0].TransactionId,
		Listmenu: listMenu,
		TransactionDate: transactionTemp[0].TransactionDate,
		TransactionTotalPrice:totalPrice,
	}
	return &transaction , nil
}

func BundleListMenuTransaction(transactionTemp []models.TransactionResponseTemp) ([]*models.TransactionMenuResponse, int,string) {
	transactionMenu := []*models.TransactionMenuResponse{}
	listAdditional := []models.AdditionalMenu{}
	var transactionTotalPrice int
	var transactionDate string
	for _,val := range transactionTemp {
		additionalPrice := 0
		for _,val2 := range val.Menu.Additional{
			listAdditional = append(listAdditional,models.AdditionalMenu{val2.AdditionalID,val2.AdditionalName,val2.AdditionalPrice})
			additionalPrice += val2.AdditionalPrice
		}
		transactionMenu = append(transactionMenu,&models.TransactionMenuResponse{
			MenuId: val.Menu.MenuId,
			MenuName: val.Menu.MenuName,
			Quantity: val.Menu.Quantity,
			MenuPrice: val.Menu.MenuPrice,
			TotalPrice: val.Menu.TotalPrice,
			Additional: listAdditional,
		})
		transactionTotalPrice = transactionTotalPrice + val.Menu.TotalPrice+ additionalPrice
		transactionDate = val.TransactionDate
	}
	return transactionMenu,transactionTotalPrice,transactionDate
}
