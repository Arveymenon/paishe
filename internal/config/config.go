package config

import (
    "log"
    "github.com/spf13/viper"
)

type Config struct {
    Port string
}

func LoadConfig() *Config {
    viper.SetConfigFile(".env")
    viper.AutomaticEnv()

    err := viper.ReadInConfig()
    if err != nil {
        log.Fatalf("Error reading config file: %v", err)
    }

    return &Config{
        Port: viper.GetString("PORT"),
    }
}
