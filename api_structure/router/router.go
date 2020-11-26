package router

import (
	"github.com/gin-gonic/gin"
	"github.com/study/api_structure/controller"
	"github.com/study/api_structure/middlewares"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Middlewares
	router.Use(middlewares.ErrorHandler)
	router.Use(middlewares.CORSMiddleware())

	//routes
	router.GET("/ping", controller.Pong)
	router.POST("/register", controller.Create)
	router.POST("/login", controller.Login)
	router.GET("/session", controller.Session)
	router.POST("/createReset", controller.InitiatePasswordReset)
	router.POST("/resetPassword", controller.ResetPassword)
	return router
}
