package routes

import (
	controllers "github.com/ricky-lopes/api/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	personController := controllers.NewPersonControler()

	v1 := router.Group("v1")
	{
		v1.GET("/person", personController.FindAll)
		v1.POST("/person", personController.Create)
		v1.DELETE("/person/:id", personController.Delete)
	}

	return v1
}
