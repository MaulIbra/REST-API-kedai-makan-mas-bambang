package api_master

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/maulIbra/clean-architecture-go/api-master/domain/menu"
)

func Init(router *mux.Router, db *sql.DB) {
	//menu
	menuRepo := menu.NewMenuRepo(db)
	menuUsecase := menu.NewMenuUsecase(menuRepo)
	menuController := menu.NewMenuController(menuUsecase)
	menuController.Menu(router)
}
