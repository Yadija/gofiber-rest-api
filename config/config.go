package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Config(key string) string {
	// load .env file
	config := viper.New()
	config.SetConfigFile(".env")
	err := config.ReadInConfig()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return config.GetString(key)
}
