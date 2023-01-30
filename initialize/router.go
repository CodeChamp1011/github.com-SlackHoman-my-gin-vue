package initialize

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/SlackHoman/my-gin-vue/server"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	PublicRouter := Router.Group("")
	{
		PublicRouter.GET("/test", func(c *gin.Context) {
			c.JSON(200, "Test sucessfully!")
		})
	}
}