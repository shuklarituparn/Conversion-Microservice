package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shuklarituparn/Conversion-Microservice/internal/email"
	"github.com/shuklarituparn/Conversion-Microservice/internal/handlers"
	"github.com/shuklarituparn/Conversion-Microservice/internal/logger"
	"github.com/shuklarituparn/Conversion-Microservice/middlewares"
	"io"
	"os"
)

func main() {

	// disabling the logs on the console
	gin.DisableConsoleColor()

	// Creating the logger for our gin router
	logs := logger.InitLog()
	defer func(logs *os.File) {
		err := logs.Close()
		if err != nil {

		}
	}(logs)

	gin.DefaultWriter = io.MultiWriter(logs)

	//creating the gin router
	router := gin.Default()

	//Using the Jaegar tracing
	router.Use(middlewares.TracingMiddleware())

	go func() {
		email.GenerateVerficationEmailConsumer()
	}()
	go func() {
		email.SendEmailConsumer()
	}()
	go func() {
		email.GenerateRestoreEmailConsumer()
	}()
	router.LoadHTMLGlob("../../templates/*")
	router.Static("/static", "../../static")
	router.Static("/uploads", "../../uploads")

	//Adding the metrics
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	//Describing all the handlers
	router.GET("/", handlers.WelcomeHandler)
	router.GET("/login", handlers.Login)
	router.GET("/callback", handlers.Callback)
	router.NoRoute(handlers.NoRouteHandler)

	protected := router.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/dashboard", handlers.Dashboard)
		protected.GET("/convert", handlers.Convert)
		protected.POST("/convert", handlers.ConvertUpload)
		protected.GET("/cut", handlers.Cut)
		protected.POST("/cut_edit", handlers.CutEditPage)
		protected.GET("/watermark", handlers.Watermark)
		protected.GET("/screenshot", handlers.Screenshot)
		protected.GET("/profile", handlers.Profile)
		protected.GET("/profile/email", handlers.EmailHandler)
		protected.POST("/profile/email", handlers.EmailUpdateHandler)
		protected.GET("/profile/files", handlers.FileHistory)
		protected.GET("/profile/delete", handlers.AccountDeleteConf)
		protected.POST("/profile/delete", handlers.AccountDelete)
		protected.GET("/profile/restore", handlers.Restore)
		protected.POST("/verify_mail", handlers.EmailConfirm)
		protected.GET("/verify_mail", handlers.VerificationEmail)

		protected.GET("/signout", handlers.Signout)
	}

	err := router.Run(":8085")
	if err != nil {
		return
	}
}
