package main

import (
	"fmt"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"io"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shuklarituparn/Conversion-Microservice/internal/database_file"
	"github.com/shuklarituparn/Conversion-Microservice/internal/email"
	"github.com/shuklarituparn/Conversion-Microservice/internal/ffmpeg"
	"github.com/shuklarituparn/Conversion-Microservice/internal/handlers"
	"github.com/shuklarituparn/Conversion-Microservice/internal/logger"
	"github.com/shuklarituparn/Conversion-Microservice/middlewares"
)

func main() {

	gin.DisableConsoleColor()

	logs := logger.InitLog()
	defer func(logs *os.File) {
		err := logs.Close()
		if err != nil {

		}
	}(logs)

	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v", err)
	}

	gin.DefaultWriter = io.MultiWriter(logs)

	router := gin.Default()

	router.Use(sentrygin.New(sentrygin.Options{}))

	go email.GenerateVerficationEmailConsumer()
	go email.SendEmailConsumer()
	go email.GenerateRestoreEmailConsumer()
	go email.DownloadMailConsumer()

	// Video Conversion
	go ffmpeg.VideoConversionConsumer()
	go ffmpeg.VideoCutConsumer()
	go ffmpeg.ScreenshotConsumer()
	go ffmpeg.WatermarkConsumer()

	// Database Operations
	go database_file.MongoUploadConsumer()
	go database_file.MongoUploadScreenshotConsumer()

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
		protected.POST("/cut", handlers.CutEditResult)
		protected.GET("/watermark", handlers.Watermark)
		protected.POST("/watermark", handlers.WatermarkResult)
		protected.GET("/screenshot", handlers.Screenshot)
		protected.POST("/screenshot_edit", handlers.Screenshot_edit)
		protected.POST("/screenshot", handlers.ScreenShotResult)
		protected.GET("/profile", handlers.Profile)
		protected.GET("/profile/email", handlers.EmailHandler)
		protected.POST("/profile/email", handlers.EmailUpdateHandler)
		protected.GET("/profile/files", handlers.FileHistory)
		protected.GET("/profile/delete", handlers.AccountDeleteConf)
		protected.POST("/profile/delete", handlers.AccountDelete)
		protected.GET("/profile/restore", handlers.Restore)
		protected.POST("/verify_mail", handlers.EmailConfirm)
		protected.GET("/profile/download", handlers.Download)
		protected.GET("/verify_mail", handlers.VerificationEmail)

		protected.GET("/signout", handlers.Signout)
	}

	err := router.Run(":8085")
	if err != nil {
		return
	}
}
