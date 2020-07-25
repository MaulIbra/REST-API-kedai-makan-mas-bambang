package category

import (
	"github.com/gorilla/mux"
	"github.com/maulIbra/clean-architecture-go/api-master/middleware"
	"github.com/maulIbra/clean-architecture-go/utils"
	"log"
	"net/http"
)

type categoryController struct {
	categoryUsecase ICategoryUsecase
}

func NewCategoryController(usecase ICategoryUsecase) *categoryController {
	return &categoryController{
		categoryUsecase: usecase,
	}
}


func (cc *categoryController) Category(r *mux.Router) {
	menu := r.PathPrefix("/category").Subrouter()
	menu.Use(middleware.TokenValidationMiddleware)
	menu.HandleFunc("", cc.readCategory).Methods(http.MethodGet)
}

func (cc *categoryController) readCategory(w http.ResponseWriter, r *http.Request)  {
	categoryList, err := cc.categoryUsecase.GetCategory()
	if err != nil {
		log.Print(err)
		utils.HandleRequest(w, http.StatusBadGateway)
	}else {
		utils.HandleResponse(w, http.StatusOK, categoryList)
	}
}

