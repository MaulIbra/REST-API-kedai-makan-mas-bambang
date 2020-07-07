package main

import (
	api_master "github.com/maulIbra/clean-architecture-go/api-master"
	"github.com/maulIbra/clean-architecture-go/infrastructure"
	"github.com/maulIbra/clean-architecture-go/infrastructure/config"
)

func main(){
	env := config.NewEnv()
	db := infrastructure.InitDB(env)
	router := infrastructure.MuxRouter()
	api_master.Init(router,db)
	infrastructure.ListenServe(router)
}
