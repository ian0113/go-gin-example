package routes

import (
	"github.com/ian0113/go-gin-example/controllers"
	"github.com/ian0113/go-gin-example/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	registerApiRouter(func(rg *gin.RouterGroup) {
		authMiddleware := middlewares.NewAuthMiddleware()
		controller := controllers.NewAuthController()
		r := rg.Group("/auth")
		r.POST("/login", controller.Login)
		r.POST("/logout", authMiddleware.ValidAuthStatus, controller.Logout)
		r.POST("/refresh", controller.Refresh)
	})
}
