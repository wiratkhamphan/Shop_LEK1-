// shoplek/shoplek.go
package shoplek

import (
	PostRegister "SHOP_LEK/API/Shop/Shop_LEK/Login/Register_user/PostRegister"
	"SHOP_LEK/database/condb"
	"log"

	"github.com/gin-gonic/gin"
)

func RunLocalhost() {
	dbInstance, err := condb.NewConndb_mysql()
	if err != nil {
		panic(err.Error())
	}
	PostRegister.DB = dbInstance
	log.Printf("Failed to insert data into the database. Error: %v\n", err)

	router := gin.Default()

	// Use RegisterUser as a gin.HandlerFunc
	router.POST("/Register_user", PostRegister.RegisterUser)

	router.Run("localhost:8080")
}
