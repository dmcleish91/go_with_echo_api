package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Using JSON decoder (recommended)
func AddDog(c echo.Context) error {
	dog := Animal{}

	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&dog)
	if err != nil {
		log.Printf("Failed procceing addDog request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("this is your dog: %#v", dog)
	return c.String(http.StatusOK, "We got your dog!")
}
