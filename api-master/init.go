package api_master

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/maulIbra/clean-architecture-go/api-master/domain/authentication"
	"github.com/maulIbra/clean-architecture-go/api-master/domain/category"
	"github.com/maulIbra/clean-architecture-go/api-master/domain/menu"
	"github.com/maulIbra/clean-architecture-go/api-master/domain/transaction"
)

func Init(router *mux.Router, db *sql.DB) {
	//menu
	menuRepo := menu.NewMenuRepo(db)
	menuUsecase := menu.NewMenuUsecase(menuRepo)
	menuController := menu.NewMenuController(menuUsecase)
	menuController.Menu(router)

	//category
	categoryRepo := category.NewCategoryRepoImpl(db)
	categoryUsecase := category.NewCategoryUsecaseImpl(categoryRepo)
	categoryController := category.NewCategoryController(categoryUsecase)
	categoryController.Category(router)

	//transaction
	transactionRepo := transaction.NewTransactionRepo(db)
	transactionUsecase := transaction.NewTransactionUsecase(transactionRepo)
	transactionController := transaction.NewTransactionController(transactionUsecase)
	transactionController.Transaction(router)

	//authentication
	AuthenticationRepo := authentication.NewAuthenticationRepo(db)
	AuthenticationUsecase := authentication.NewAuthenticationUsecase(AuthenticationRepo)
	AuthenticationController := authentication.NewAuthenticationController(AuthenticationUsecase)
	AuthenticationController.Authenticate(router)
}
