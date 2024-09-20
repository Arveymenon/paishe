package config

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/viper"
	// "gorm.io/driver/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// "github.com/Arveymenon/paishe/internal/repository"
	"github.com/Arveymenon/paishe/internal/domain"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

var DB *gorm.DB

func GetDatabaseConfig() *DatabaseConfig {
	port, err := strconv.Atoi(viper.GetString("DB_PORT"))
	if err != nil {
		log.Printf("Invalid port id provided %d", port)
	}
	return &DatabaseConfig{
		Host:     viper.GetString("DB_HOST"),
		Port:     port,
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASSWORD"),
		DBName:   viper.GetString("DB_NAME"),
	}
}

func SetUpDataBase() *gorm.DB {
	dbConfig := GetDatabaseConfig()

	log.Println("Connecting Database on port", dbConfig.Port)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	// Create tables if not created in the system
	db.AutoMigrate(&domain.User{})

	return db
}
