package shoplek

import (
	"log"

	PostRegister "SHOP_LEK/API/Shop/Shop_LEK/Login/Register_user/PostRegister"
	condb "SHOP_LEK/database/condb"

	"github.com/gin-gonic/gin"
)

func RunLocalhost() {
	var err error
	_, err = condb.NewConndb_mysql()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.POST("/Register_user", PostRegister.PostAlbums)

	router.Run("localhost:8080")
}
