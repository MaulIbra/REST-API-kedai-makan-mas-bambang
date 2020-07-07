package config

import "github.com/maulIbra/clean-architecture-go/utils"

type Env struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	SchemaName string
	Driver string
}

func NewEnv() *Env {
	return &Env{
		DbUser:     utils.GetEnv("dbUser", "defaultUser"),
		DbPassword: utils.GetEnv("dbPassword", "defaultPass"),
		DbHost:     utils.GetEnv("dbHost", "defaultHost"),
		DbPort:     utils.GetEnv("dbPort", "12345"),
		SchemaName: utils.GetEnv("dbSchema", "defaultScheme"),
		Driver: utils.GetEnv("driver","mysql"),
	}
}
