package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func HandleLogin(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	// check username and password against DB after hashing the password
	if username != "username" && password != "password" {
		return c.String(http.StatusUnauthorized, "Credential Incorrect")
	}

	//cookie := &http.Cookie{}
	cookie := new(http.Cookie)
	cookie.Name = "sessionID"
	cookie.Value = "some_string"
	cookie.Expires = time.Now().Add(48 * time.Hour)

	c.SetCookie(cookie)

	token, err := createJwtToken()
	if err != nil {
		log.Println("Error Creating JWT token", err)
		return c.String(http.StatusInternalServerError, "something went wrong")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "You were logged in!",
		"token":   token,
	})

}

func createJwtToken() (string, error) {
	claims := JwtClaims{
		"jack",
		jwt.StandardClaims{
			Id:        "main_user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rawToken.SignedString([]byte("mySecret"))
	if err != nil {
		return "nil", err
	}

	return token, nil
}
