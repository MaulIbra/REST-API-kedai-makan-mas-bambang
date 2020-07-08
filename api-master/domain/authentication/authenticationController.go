package authentication

import (
	"github.com/gorilla/mux"
	"github.com/maulIbra/clean-architecture-go/utils"
	"log"
	"net/http"
)

func Authenticate(r *mux.Router){
	r.HandleFunc("/token",GetToken)
}

func GetToken(w http.ResponseWriter, r *http.Request){
	token, err := utils.JwtEncoder("maulana","Rahasia dong")
	if err != nil {
		log.Print(err)
		utils.HandleRequest(w, http.StatusBadRequest)
	}else {
		utils.HandleResponse(w, http.StatusOK, token)
	}
}
