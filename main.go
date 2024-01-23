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
		// Add more protected routes here
		// protected.GET("/profile", handlers.Profile)
		// protected.GET("/settings", handlers.Settings)
	}

	router.Run(":8085")
}

//TODO: Where to use the database to store User data and sessions?

//localhost:8085/?videoaction=convert gives URL query  map[videoAction:[Convert]]
