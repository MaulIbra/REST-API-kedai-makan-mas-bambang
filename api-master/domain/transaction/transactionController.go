package transaction

import (
	"github.com/gorilla/mux"
	"github.com/maulIbra/clean-architecture-go/utils"
	"log"
	"net/http"
)

type transactionController struct {
	usecase ITransactionUsecase
}

func NewTransactionController(usecase ITransactionUsecase) *transactionController{
	return &transactionController{usecase: usecase}
}

func (th *transactionController) Transaction(r *mux.Router) {
	r.HandleFunc("/transaction", th.readTransaction).Methods(http.MethodGet)
	r.HandleFunc("/transaction/daily", th.readTransactionDaily).Methods(http.MethodGet)
	r.HandleFunc("/transaction", th.addTransaction).Methods(http.MethodPost)
}

func (th *transactionController) readTransaction(w http.ResponseWriter, r *http.Request){
	transactionList, err := th.usecase.GetTransaction("%")
	if err != nil {
		log.Print(err)
		utils.HandleRequest(w, http.StatusBadGateway)
	}else {
		utils.HandleResponse(w, http.StatusOK, transactionList)
	}
}

func (th *transactionController) readTransactionDaily(w http.ResponseWriter, r *http.Request){
	date := utils.DecodeQueryParams("date",r)
	transactionList, err := th.usecase.GetTransaction(date+"%")
	var totalOmset int
	for _,val := range transactionList{
		totalOmset += val.TransactionTotalPrice
	}
	transactionOmset := TransactionOmset{
		TotalOmset: totalOmset,
		ListTransaction: transactionList,
	}
	if err != nil {
		utils.HandleRequest(w, http.StatusBadGateway)
	}else{
		utils.HandleResponse(w, http.StatusOK, transactionOmset)
	}
}

func (th *transactionController) addTransaction(w http.ResponseWriter, r *http.Request){
	var transaction Transaction
	err := utils.JsonDecoder(&transaction,r)
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	}
	err = th.usecase.PostTransaction(&transaction)
	if err != nil {
		log.Print(err)
		utils.HandleRequest(w, http.StatusBadGateway)
	}else{
		transaction,err := th.usecase.GetTransactionByID(transaction.TransactionId)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		}else{
			utils.HandleResponse(w, http.StatusOK,transaction)
		}
	}
}

