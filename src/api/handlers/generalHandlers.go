package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Person struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func HandleHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from a golang API!")
}

func GetPerson(c echo.Context) error {
	personName := c.QueryParam("name")
	personAge := c.QueryParam("age")

	dataType := c.Param("data")

	if dataType == "string" {
		response := "Your name is: %s and you are %s years old."
		return c.String(http.StatusOK, fmt.Sprintf(response, personName, personAge))
	}

	if dataType == "json" {
		person := Person{
			Name: personName,
			Age:  personAge,
		}

		return c.JSON(http.StatusOK, person)
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "Please define data as 'string' or 'json'",
	})
}
