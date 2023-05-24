package main

import (
	"log"
	"os"

	"github.com/devfurkankizmaz/go-lib-management-app/configs"
	"github.com/devfurkankizmaz/go-lib-management-app/models"
)

func init() {
	config, err := configs.NewEnv()
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	configs.NewDBConnection(config)
}

func main() {
	err := configs.App().DB.AutoMigrate(&models.User{}, &models.Book{})
	if err != nil {
		log.Fatal("Migration Failed:  \n", err.Error())
		os.Exit(1)
	}
}
