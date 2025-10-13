package config

import (
	"fmt"
	"go-gin-album/pkg/util"
	"os"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
	TimeZone string
}

func (c *DBConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		c.Host, c.User, c.Password, c.Name, c.Port, c.SSLMode, c.TimeZone,
	)
}

func LoadDBConfig() *DBConfig {
	util.LoadEnv()
	dbConfig := DBConfig{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		SSLMode:  "disable",
		TimeZone: "Asia/Singapore",
	}

	return &dbConfig
}
