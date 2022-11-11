package routes

import (
	"github.com/ian0113/go-gin-mvc-example/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	registerApiRouter(func(rg *gin.RouterGroup) {
		controller := controllers.NewOrderController()
		r := rg.Group("/orders")
		r.POST("", controller.CreateOrder)
		r.GET("", controller.ListOrders)
		r.GET("/:id", controller.GetOrder)
		r.PUT("/:id", controller.UpdateOrder)
		r.DELETE("/:id", controller.DeleteOrder)
	})
}
