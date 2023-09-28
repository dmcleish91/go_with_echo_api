package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Animal struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// Most performant but very verbose
func AddCat(c echo.Context) error {
	cat := Animal{}

	defer c.Request().Body.Close()
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body for addCat: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &cat)
	if err != nil {
		log.Printf("Failed Unmarshal in addCat: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("this is your cat: %#v", cat)
	return c.String(http.StatusOK, "We got your cat!")
}

// getting query params
func GetCat(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	response := "Your cat's name is: %s\nand his type is: %s"
	return c.String(http.StatusOK, fmt.Sprintf(response, catName, catType))
}
