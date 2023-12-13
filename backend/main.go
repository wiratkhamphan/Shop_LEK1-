package main

import (
	Env "SHOP_LEK/database/Env"
	condb "SHOP_LEK/database/condb"
)

func main() {
	Env.NewEnv()
	condb.NewConndb_mysql()

}
