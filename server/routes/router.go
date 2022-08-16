package routes

import (
	"github.com/estudo-go/Desafio-cbt/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		transactions := main.Group("trasactions")
		{
			transactions.GET("/", controllers.ShowTransaction)
		}
	}

	return router
}
