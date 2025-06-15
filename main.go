package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sb-67-go-quiz-salman-input-data-buku/config"
	"github.com/sb-67-go-quiz-salman-input-data-buku/database/connection"
	_ "github.com/sb-67-go-quiz-salman-input-data-buku/docs"
	api "github.com/sb-67-go-quiz-salman-input-data-buku/router"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// init config
	config.Init()

	// init database
	connection.Init()

	// init web
	web := gin.Default()

	// route
	api.ListAllRouter(web)

	web.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// run web
	web.Run(viper.GetString("App.Port"))
}
