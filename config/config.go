package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigFile("./config/config.json")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return
	}

	env := viper.GetString("App.Env")

	_ = viper.BindEnv("App.Env", "APP_ENV")
	if overridden := viper.GetString("App.Env"); overridden != "" {
		env = overridden
	}

	fmt.Println("Current Env:", env)

	// persiapan override variable database production
	
}
