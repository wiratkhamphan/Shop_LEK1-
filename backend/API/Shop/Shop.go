package shop

import (
	"github.com/gin-gonic/gin"
)

func RunLocalhost() {

	router := gin.Default()
	// router.POST("/Register_user", PostRegister.PostAlbums)

	router.Run("localhost:8080")
}
