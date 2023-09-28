package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MainJwt(c echo.Context) error {
	user := c.Get("user")

	log.Println(user)
	return c.String(http.StatusOK, "You are on the secret jwt page")
}
