package main

import (
	"net/http"

	"github.com/devfurkankizmaz/go-lib-management-app/api/routes"
	"github.com/devfurkankizmaz/go-lib-management-app/configs"
	_ "github.com/devfurkankizmaz/go-lib-management-app/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Echo Swagger Example API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http

func main() {
	// Echo Instance
	server := echo.New()

	// Middleware
	//server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.CORS())
	app := configs.App()
	config := app.Env
	routes.Setup(config, app.DB, server)

	server.GET("/", HealthCheck)
	server.GET("/swagger/*", echoSwagger.WrapHandler)
	server.Logger.Fatal(server.Start(config.ServerPort))
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "It's Healthy!")
}
