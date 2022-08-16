package routes

import "github.com/gin-gonic/gin"

func configRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		transactions := main.Group("trasactions")
		{
			transactions.GET("/")
		}
	}
	return router
}
