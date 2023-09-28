package routes

import (
	"github.com/dmcleish91/go_echo_api/src/api/handlers"
	"github.com/labstack/echo/v4"
)

func SetMainRoutes(e *echo.Echo) {
	e.GET("/cats", handlers.GetCat)
	e.POST("/cats", handlers.AddCat)

	e.GET("/", handlers.HandleHello)
	e.GET("/login", handlers.HandleLogin)

	e.POST("/dogs", handlers.AddDog)
	e.POST("/hamsters", handlers.AddHamster)
}
