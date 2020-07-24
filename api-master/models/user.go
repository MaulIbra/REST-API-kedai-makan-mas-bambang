package models

type User struct {
	UserID   string `json:"userId"`
	Username string	`json:"username"`
	Password string	`json:"password"`
	Token Token		`json:"authentication"`
}
