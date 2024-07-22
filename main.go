package main

import (
	"ego/app"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"

	_ "ego/docs"
)

//	@title			Swagger ECHO Go API
//	@version		1.0
//	@description	This is a sample ECHO golang API.
//	@termsOfService	http://swagger.io/terms/

// @BasePath	/
func main() {
	logger := app.InitLogger()
	loggerConfig := app.InitLoggerConfig()

	conf, err := app.InitConfig()
	if err != nil {
		logger.Error(err.Error())
		panic("")
	}

	logger.With("env", conf.GetString("env"))

	e := app.Init(conf, logger, loggerConfig)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(":4242"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
