package models

type Transaction struct {
	TransactionId string              `json:transactionId`
	ListMenu []TransactionMenuRequest `json:"listMenu"`
	TransactionDate string            `json:"transactionDate"`
}

type TransactionMenuRequest struct {
	MenuId string `json:"menuId"`
	Quantity int `json:"quantity"`
	Additional []AdditionalMenu `json:"additional"`
}

type TransactionResponse struct {
	TransactionId string                `json:transactionId`
	Listmenu []*TransactionMenuResponse `json:"listmenu"`
	TransactionDate string              `json:"transactionDate"`
	TransactionTotalPrice int           `json:"transactionTotalPrice"`
}

type TransactionResponseTemp struct {
	TransactionId   string
	Menu            TransactionMenuResponse
	TransactionDate string
}

type TransactionOmset struct {
	TotalOmset int                         `json:"totalOmset"`
	ListTransaction []*TransactionResponse `json:"listTransaction"`
}

type TransactionMenuResponse struct {
	MenuId string `json:"menuId"`
	MenuName string `json:menuName`
	Quantity int `json:"quantity"`
	MenuPrice int `json:"menuPrice"`
	TotalPrice int `json:"totalPrice"`
	Additional []AdditionalMenu `json:"additional"`
}

