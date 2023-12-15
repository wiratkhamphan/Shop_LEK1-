// main.go or wherever you initialize your application
package main

import (
	RunLocalhost "SHOP_LEK/API/Shop/Shop_LEK"
	condb "SHOP_LEK/database/condb"
)

func main() {

	// // ... other initialization code
	// // Run your application
	condb.NewConndb_mysql()
	RunLocalhost.RunLocalhost()
}
