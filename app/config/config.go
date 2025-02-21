package config

import (
	"fmt"
	"github.com/saadrupai/order-assignment/app/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port       string `json:"port"`
	DbUser     string `json:"db_user"`
	DbPassword string `json:"db_password"`
	DbHost     string `json:"db_host"`
	DbPort     string `json:"db_port"`
	DbSchema   string `json:"db_schema"`
}

var LocalConfig *Config
var db *gorm.DB

func LoadConfig() *Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("failed to load env variables")
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("failed to load env variables")
	}

	ConnectDB(config)

	return &config
}

func ConnectDB(config Config) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DbSchema)
	gormDB, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		log.Fatal("failed to connect to database")
	}

	gormDB.AutoMigrate(
		&entity.User{},
		&entity.Order{})

	db = gormDB
}

func GetDb() *gorm.DB {
	return db
}

func SetConfig() {
	LocalConfig = LoadConfig()
}
