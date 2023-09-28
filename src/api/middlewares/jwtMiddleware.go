package middlewares

import (
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func SetJwtMiddlewares(g *echo.Group) {
	g.Use(echojwt.JWT([]byte("mySecret")))
}
