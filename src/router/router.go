package router

import (
	"github.com/dmcleish91/go_echo_api/src/api/middlewares"
	"github.com/dmcleish91/go_echo_api/src/api/routes"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	//create groups
	adminGroup := e.Group("/admin")
	cookieGroup := e.Group("/cookie")
	jwtGroup := e.Group("/jwt")

	middlewares.SetMainMiddlewares(e)
	middlewares.SetAdminMiddlewares(adminGroup)
	middlewares.SetCookieMiddlewares(cookieGroup)
	middlewares.SetJwtMiddlewares(jwtGroup)

	routes.SetMainRoutes(e)
	routes.SetAdminRoutes(adminGroup)
	routes.SetCookieRoutes(cookieGroup)
	routes.SetJwtRoutes(jwtGroup)

	return e
}
