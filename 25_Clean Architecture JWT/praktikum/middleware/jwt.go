package middleware

import (
	"belajar-go-echo/model"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

const secretJwt = "d6ca6c89-ce38-49d4-8d8e-db0c553b4f25"

func CreateJwtToken(username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["username"] = username
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretJwt))
}

func LoginHandler(c echo.Context) error {
	loginData := model.User{}
	c.Bind(&loginData)
	// username: admin, password: rahasia
	if loginData.Username == "admin" && loginData.Password == "rahasia" {
		token, err := CreateJwtToken(loginData.Username)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Cannot create token")
		}
		return c.JSON(http.StatusOK, map[string]string{
			"success": "ok",
			"token":   token,
		})
	}
	return c.String(http.StatusUnauthorized, "Unauthorized")
}
