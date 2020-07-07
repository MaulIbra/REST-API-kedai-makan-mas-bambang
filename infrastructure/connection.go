package infrastructure

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/maulIbra/clean-architecture-go/infrastructure/config"
)

func InitDB(env *config.Env) (*sql.DB) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", env.DbUser, env.DbPassword, env.DbHost, env.DbPort, env.SchemaName)
	db, _ := sql.Open(env.Driver, dataSource)

	if err := db.Ping(); err != nil {
		panic(err)
		return nil
	}
	return db
}


