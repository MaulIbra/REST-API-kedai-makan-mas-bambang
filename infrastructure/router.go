package infrastructure

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/maulIbra/clean-architecture-go/utils"
	"net/http"
)

func MuxRouter() *mux.Router{
	return mux.NewRouter()
}

func ListenServe(router *mux.Router){
	port := utils.GetEnv("serverPort","3000")
	host := utils.GetEnv("serverHost","localhost")
	err := http.ListenAndServe(fmt.Sprintf("%v:%v",host,port),router)
	if err != nil {
		panic(err)
	}
}