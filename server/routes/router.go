package routes

import (
	"github.com/estudo-go/Desafio-cbt/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		transactions := main.Group("transacao")
		{
			transactions.GET("/:tipo", controllers.ShowTransaction)
			transactions.GET("/", controllers.ShowTransaction)
			//transactions.GET("/", controllers.)
		}
	}

	return router
}
