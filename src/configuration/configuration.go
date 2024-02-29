package configuration

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Configuration struct {
	Database struct {
		Host     string
		User     string
		Password string
		DBName   string
		Port     int
		SSL      string
	}
	System struct {
		Port string
	}
}

type DevConfiguration struct {
	Database struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
		Port     int    `json:"db_port"`
		SSL      string `json:"ssl"`
	} `json:"database"`

	System struct {
		Port string `json:"system_port"`
	} `json:"system"`
}

func LoadEnvironmentVars() (Configuration, error) {
	environment := os.Getenv("PROJECT_ENV")

	var config Configuration

	if len(environment) == 0 {
		err := errors.New("Environment var not defined, please set PROJECT_ENV in your system's environment variables")

		return config, err
	}

	if environment == "development" {
		err := loadDevEnvironmentVars(&config)

		return config, err
	}

	loadEnvironment(&config)

	return config, nil
}

func loadDevEnvironmentVars(config *Configuration) error {
	file, err := os.Open("env.json")
	if err != nil {
		fmt.Println("Error opening config file:", err)

		return err
	}
	defer file.Close()

	var devConfig DevConfiguration

	err = json.NewDecoder(file).Decode(&devConfig)

	if err != nil {
		fmt.Println("Error decoding config dev env:", err)
		return err
	}

	config.Database.Host = devConfig.Database.Host
	config.Database.Password = devConfig.Database.Password
	config.Database.DBName = devConfig.Database.DBName
	config.Database.Port = devConfig.Database.Port
	config.Database.User = devConfig.Database.User
	config.Database.SSL = devConfig.Database.SSL
	config.System.Port = devConfig.System.Port

	return nil
}

func loadEnvironment(config *Configuration) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	ssl := os.Getenv("DB_SSL")

	port := os.Getenv("SYSTEM_PORT")

	parsedDbPort, _ := strconv.Atoi(dbPort)

	config.Database.Host = host
	config.Database.User = user
	config.Database.Password = password
	config.Database.DBName = dbname
	config.Database.Port = parsedDbPort
	config.Database.SSL = ssl
	config.System.Port = port
}
