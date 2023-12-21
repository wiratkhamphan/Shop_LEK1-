package condb

import (
	Env "SHOP_LEK/database/Env"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewConndb_mysql() (*sql.DB, error) {
	env, err := Env.NewEnv()
	if err != nil {
		return nil, err
	}

	// db, err := sql.Open(env.DBDriver, env.DBUser+":"+env.DBPass+"@/"+env.DBName)
	db, err := sql.Open(env.DBDriver, fmt.Sprintf("%s:%s@/%s", env.DBUser, env.DBPass, env.DBName))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	// defer db.Close()
	return db, err

}
