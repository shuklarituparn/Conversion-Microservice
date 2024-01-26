package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shuklarituparn/Conversion-Microservice/handlers"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	router.GET("/", handlers.WelcomeHandler)
	router.GET("/login", handlers.Login)
	router.GET("/callback", handlers.Callback)
	protected := router.Group("/")

	protected.Use(handlers.AuthMiddleware())
	{
		protected.GET("/dashboard", handlers.Dashboard)
		protected.GET("/convert", handlers.Dashboard)
		protected.GET("/cut", handlers.Dashboard)
		protected.GET("/watermark", handlers.Dashboard)
		protected.GET("/extract", handlers.Dashboard)
		protected.GET("/profile", handlers.Dashboard)
		protected.GET("/signout", handlers.Dashboard)
	}

	router.Run(":8085")
}

//TODO: Where to use the database to store User data and sessions?
