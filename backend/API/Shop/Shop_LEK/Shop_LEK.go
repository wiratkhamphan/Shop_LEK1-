package shoplek

import (
	"log"
	"net/http"

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
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	router.POST("/Register_user", PostRegister.PostAlbums)

	router.Run("localhost:8080")
}
