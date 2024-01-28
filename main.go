package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shuklarituparn/Conversion-Microservice/handlers"
	"io"
	"os"
)

func main() {

	gin.DisableConsoleColor()

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", handlers.WelcomeHandler)
	router.GET("/login", handlers.Login)
	router.GET("/callback", handlers.Callback)
	router.NoRoute(handlers.NoRouteHandler)
	protected := router.Group("/")

	protected.Use(handlers.AuthMiddleware())
	{
		protected.GET("/dashboard", handlers.Dashboard)
		protected.GET("/convert", handlers.Convert)
		protected.GET("/cut", handlers.Cut)
		protected.GET("/watermark", handlers.Watermark)
		protected.GET("/extract", handlers.Extract)
		protected.GET("/profile", handlers.Profile)
		protected.GET("/signout", handlers.Signout)
	}

	router.Run(":8085")
}
