package menu

import (
	"github.com/gorilla/mux"
	"github.com/maulIbra/clean-architecture-go/utils"
	"log"
	"net/http"
)

type menuController struct {
	usecase IMenuUsecase
}

func NewMenuController(usecase IMenuUsecase) *menuController {
	return &menuController{usecase: usecase}
}

func (ph *menuController) Menu(r *mux.Router) {
	r.HandleFunc("/menu", ph.readMenu).Methods(http.MethodGet)
	r.HandleFunc("/menu/{id}", ph.readMenuById).Methods(http.MethodGet)
	r.HandleFunc("/menu", ph.addMenu).Methods(http.MethodPost)
	r.HandleFunc("/menu/{id}", ph.editMenu).Methods(http.MethodPut)
	r.HandleFunc("/menu/{id}", ph.deleteMenu).Methods(http.MethodDelete)

}

func (ph *menuController) readMenu(res http.ResponseWriter, req *http.Request) {
	menuList, err := ph.usecase.GetMenu()
	if err != nil {
		log.Print(err)
		utils.HandleRequest(res, http.StatusBadGateway)
	}else {
		utils.HandleResponse(res, http.StatusOK, menuList)
	}
}

func (ph *menuController) readMenuById(res http.ResponseWriter, req *http.Request) {
	id := utils.DecodePathVariabel("id", req)
	menu, err := ph.usecase.GetMenuByID(id)
	var x []string
	if err != nil {
		utils.HandleResponse(res, http.StatusOK,x )
	} else {
		utils.HandleResponse(res, http.StatusOK, menu)
	}
}

func (ph *menuController) addMenu(res http.ResponseWriter, req *http.Request) {
	var menu Menu
	err := utils.JsonDecoder(&menu, req)
	log.Print(menu)
	if err != nil {
		utils.HandleRequest(res, http.StatusBadRequest)
	} else {
		err = ph.usecase.PostMenu(&menu)
		if err != nil {
			utils.HandleRequest(res, http.StatusBadGateway)
		} else {
			menu2,err := ph.usecase.GetMenuByID(menu.MenuiD)
			if err != nil {
				log.Print(err)
			}else{
				utils.HandleResponse(res, http.StatusCreated,menu2)
			}
		}
	}
}

func (ph *menuController) editMenu(res http.ResponseWriter, req *http.Request) {
	var menu Menu

	id := utils.DecodePathVariabel("id", req)
	err := utils.JsonDecoder(&menu, req)
	if err != nil {
		log.Print(err)
	}
	menu.MenuiD = id
	err = ph.usecase.UpdateMenu(&menu)
	if err != nil {
		utils.HandleRequest(res, http.StatusBadGateway)
	} else {
		menuUpdate, err := ph.usecase.GetMenuByID(id)
		if err != nil {
			log.Print(err)
		} else {
			utils.HandleResponse(res, http.StatusOK, menuUpdate)
		}
	}
}

func (ph *menuController) deleteMenu(res http.ResponseWriter, req *http.Request) {
	id := utils.DecodePathVariabel("id", req)
	if len(id) > 0 {
		err := ph.usecase.DeleteMenu(id)
		if err != nil {
			utils.HandleRequest(res, http.StatusNotFound)
		} else {
			utils.HandleRequest(res, http.StatusOK)
		}
	} else {
		utils.HandleRequest(res, http.StatusBadRequest)
	}
}
