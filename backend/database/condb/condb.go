package condb

import (
	Env "SHOP_LEK/database/Env"
	quan "SHOP_LEK/database/condb/quan"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewConndb_mysql() (*sql.DB, error) {
	env, err := Env.NewEnv()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(env.DBDriver, env.DBUser+":"+env.DBPass+"@/"+env.DBName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	// print(db)
	fmt.Println("Successfully connected!")
	USER, err := quan.GetUser(db, 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Quan", USER)
	return db, nil
}
