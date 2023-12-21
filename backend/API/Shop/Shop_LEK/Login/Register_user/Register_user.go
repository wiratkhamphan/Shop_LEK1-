// registeruser/registeruser.go
package registeruser

import (
	PostRegister "SHOP_LEK/API/Shop/Shop_LEK/Login/Register_user/PostRegister"

	"github.com/gin-gonic/gin"
)

func PostAlbums_1(c *gin.Context) {
	PostRegister.PostAlbums(c.Copy())
}
