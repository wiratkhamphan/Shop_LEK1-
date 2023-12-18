// main.go or wherever you initialize your application
package main

import (
	RunLocalhost "SHOP_LEK/API/Shop/Shop_LEK"
	condb "SHOP_LEK/database/condb"
	"database/sql"
)

var db *sql.DB

func main() {

	db, err := condb.NewConndb_mysql()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	RunLocalhost.RunLocalhost()
}
