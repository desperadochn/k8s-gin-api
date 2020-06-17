package router

import (
	"github.com/gin-gonic/gin"
	"jarvis_server/controllers"
)

func InitRouter() {
	router := gin.Default()

	dev := router.Group("dev-api")
	{
		user := dev.Group("user")
		{
			user.POST("login", controllers.Login)
			user.GET("info", controllers.Info)
		}

		node := dev.Group("node")
		{
			node.GET("list", controllers.NodeList)
		}

		namespaces := dev.Group("namespaces")
		{
			namespaces.GET("list", controllers.NameSpacesList)
		}
	}

	router.Run(":8080")
}
