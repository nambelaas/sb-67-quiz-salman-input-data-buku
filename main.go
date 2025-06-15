package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sb-67-go-quiz-salman-input-data-buku/config"
	"github.com/sb-67-go-quiz-salman-input-data-buku/database/connection"
	api "github.com/sb-67-go-quiz-salman-input-data-buku/router"
	"github.com/spf13/viper"
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

	// run web
	web.Run(viper.GetString("App.Port"))
}
