package condb

import (
	Env "SHOP_LEK/database/Env"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewConndb_mysql() (*sql.DB, error) {
	env, err := Env.NewEnv()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(env.DBDriver, env.DBUser+":"+env.DBPass+"@/"+env.DBName)

	return db, err

}
