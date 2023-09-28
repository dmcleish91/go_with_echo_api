package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// JSON handling - idiomatic for echo
func AddHamster(c echo.Context) error {
	hamster := Animal{}

	err := c.Bind(&hamster)
	if err != nil {
		log.Printf("Failed procceing addHamster request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("this is your hamster %#v", hamster)
	return c.String(http.StatusOK, "We got your hamster")
}
