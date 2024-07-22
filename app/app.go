package app

import (
	"ego/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
	"github.com/spf13/viper"
	"log/slog"
)

func Init(conf *viper.Viper, logger *slog.Logger, loggerConfig slogecho.Config) *echo.Echo {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(slogecho.NewWithConfig(logger, loggerConfig))
	e.Use(middleware.Recover())
	//e.Use(echojwt.WithConfig(echojwt.Config{
	//	SigningKey: []byte(conf.GetString("JwtSigningKey")),
	//}))

	sayHandler := handler.NewSay()

	// Routes
	e.GET("/hello", sayHandler.Hello)
	e.GET("/bye", sayHandler.Bye)
	e.GET("/damn", sayHandler.Damn)

	return e
}
