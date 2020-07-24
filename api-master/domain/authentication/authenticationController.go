package authentication

import (
	"github.com/gorilla/mux"
	"github.com/maulIbra/clean-architecture-go/api-master/models"
	"github.com/maulIbra/clean-architecture-go/utils"
	"log"
	"net/http"
)

type AuthenticationController struct {
	usecase IAuthenticationUsecase
}

func NewAuthenticationController(usecase IAuthenticationUsecase) *AuthenticationController {
	return &AuthenticationController{usecase: usecase}
}

func (a *AuthenticationController)Authenticate(r *mux.Router){
	user := r.PathPrefix("/user").Subrouter()
	user.HandleFunc("/login",a.Login).Methods(http.MethodPost)
	user.HandleFunc("",a.Registration).Methods(http.MethodPost)
	user.HandleFunc("",a.ReadUser).Methods(http.MethodGet)
	user.HandleFunc("/{id}",a.DeleteUser).Methods(http.MethodDelete)

}

func (a *AuthenticationController) Registration(w http.ResponseWriter, r *http.Request) {
	var profile models.Profile
	err := utils.JsonDecoder(&profile, r)
	profile.User.Password = utils.Encrypt([]byte(profile.User.Password))
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	}else {
		err = a.usecase.AddUserProfile(&profile)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		} else {
			user,err := a.usecase.ReadUserByEmail(profile.User.Username)
			if err != nil {
				log.Print(err)
			}else{
				utils.HandleResponse(w, http.StatusCreated,user)
			}
		}
	}
}

func (a *AuthenticationController) Login(w http.ResponseWriter, r *http.Request){
	var user models.User
	err := utils.JsonDecoder(&user, r)
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest,utils.BAD_REQUEST)
	}else{
		userTemp,err := a.usecase.ReadUserByEmail(user.Username)
		if err != nil {
			utils.HandleResponseError(w, http.StatusBadGateway,utils.BAD_GATEWAY)
		}
		isValid := utils.CompareEncrypt(userTemp.Password, []byte(user.Password))
		if isValid {
			token, err := utils.JwtEncoder(userTemp.Username,"Rahasia dong")
			if err != nil {
				utils.HandleResponseError(w,http.StatusBadRequest,utils.BAD_REQUEST)
			}
			userTemp.Token = models.Token{Key: token}
			utils.HandleResponse(w, http.StatusOK,userTemp)
		}else{
			utils.HandleResponseError(w,http.StatusOK,"Wrong password or username")
		}
	}
}


func (a *AuthenticationController) ReadUser(w http.ResponseWriter, r *http.Request){
	profile, err := a.usecase.ReadUser()
	if err != nil {
		utils.HandleResponseError(w,http.StatusBadRequest,utils.BAD_REQUEST)
	}else{
		utils.HandleResponse(w,http.StatusOK,profile)
	}
}

func (a *AuthenticationController) DeleteUser(w http.ResponseWriter, r *http.Request){
	id := utils.DecodePathVariabel("id", r)
	if len(id) > 0 {
		err := a.usecase.DeleteUser(id)
		if err != nil {
			utils.HandleRequest(w, http.StatusNotFound)
		} else {
			utils.HandleRequest(w, http.StatusOK)
		}
	} else {
		utils.HandleRequest(w, http.StatusBadRequest)
	}
}


