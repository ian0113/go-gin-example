package routes

import (
	"github.com/ian0113/go-gin-mvc-example/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	registerApiRouter(func(rg *gin.RouterGroup) {
		controller := controllers.NewUserController()
		r := rg.Group("/users")
		r.POST("", controller.Register)
		r.DELETE("/:id", controller.Unregister)
	})
}
