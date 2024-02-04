package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shuklarituparn/Conversion-Microservice/internal/email"
	"github.com/shuklarituparn/Conversion-Microservice/internal/handlers"

	"io"
	"os"
)

func main() {

	gin.DisableConsoleColor()

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()

	router.LoadHTMLGlob("../../templates/*")
	router.Static("/static", "../../static")
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.GET("/", handlers.WelcomeHandler)
	router.GET("/login", handlers.Login)
	router.GET("/callback", handlers.Callback)
	router.NoRoute(handlers.NoRouteHandler)
	protected := router.Group("/")

	protected.Use(handlers.AuthMiddleware())
	{
		protected.GET("/dashboard", handlers.Dashboard)
		protected.GET("/convert", handlers.Convert)
		protected.POST("/convert", handlers.ConvertUpload)
		protected.GET("/cut", handlers.Cut)
		protected.GET("/watermark", handlers.Watermark)
		protected.GET("/screenshot", handlers.Screenshot)
		protected.GET("/profile", handlers.Profile)
		protected.GET("/signout", handlers.Signout)
	}
	email.ConsumeEmail()

	router.Run(":8085")
}
