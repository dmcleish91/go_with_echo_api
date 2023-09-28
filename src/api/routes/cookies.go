package routes

import (
	"github.com/dmcleish91/go_echo_api/src/api/handlers"
	"github.com/labstack/echo/v4"
)

func SetCookieRoutes(g *echo.Group) {
	g.GET("/main", handlers.MainCookie)
}
