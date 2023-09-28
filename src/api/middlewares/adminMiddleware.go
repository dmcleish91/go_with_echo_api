package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetAdminMiddlewares(g *echo.Group) {
	// configure middleware for group
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	// configure basic authentication
	g.Use(middleware.BasicAuth(func(username, password string, ctx echo.Context) (bool, error) {
		// true condition
		if username == "gecho" && password == "gecho" {
			return true, nil
		}
		// or you can simply let echo handle it for you {"message": "Unauthorized"}
		return false, nil
	}))
}
