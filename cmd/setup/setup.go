package setup

import (
	"log"

	"github.com/book-service/api/app/config"
	"github.com/book-service/api/app/db/connection"
	"github.com/book-service/api/app/dependencies"
	"github.com/book-service/api/app/helper/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// AppSetup initializes all the necessary setups
func AppSetup() (*gin.Engine, *config.Config) {
	setUpLogger()

	// Load configuration
	configData, err := config.LoadConfigs("config/config.yaml")
	if err != nil {
		log.Fatal("failed to load config file", err)
	}
	loadEnv(configData.App.Environment)
	// Load environment variables

	// Initialize databases
	connection.MustInit()
	db := connection.DB()
	dependencies.InitializeContainer(db)

	// Initialize Gin router
	router := gin.Default()

	return router, configData
}

func loadEnv(environment string) {
	file := ".env.dev"
	if environment == "production" {
		file = ".env.prod"
	}

	if err := godotenv.Load(file); err != nil {
		log.Fatalln("Failed to load env file")
		panic(err)
	}

	log.Println("App is running on", environment, "environment")
}

func setUpLogger() {
	loggerOptions := &logger.LoggerSetUpOptions{
		Info: &logger.LoggerOptions{
			Filename:   "./logs/info.log",
			MaxSize:    5,
			MaxBackups: 2,
			MaxAge:     5,
			Compress:   true,
		},
		Error: &logger.LoggerOptions{
			Filename:   "./logs/error.log",
			MaxSize:    5,
			MaxBackups: 2,
			MaxAge:     5,
			Compress:   true,
		},
		Warn: &logger.LoggerOptions{
			Filename:   "./logs/warn.log",
			MaxSize:    5,
			MaxBackups: 2,
			MaxAge:     5,
			Compress:   true,
		},
	}
	logger.Init(loggerOptions)
}
