package router

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		userName := c.FormValue("username")
		password := c.FormValue("password")

		if userName == "jon" && password == "shhh!" {
			claims := &jwtCustomClaims{
				"Jon Snow",
				true,
				jwt.RegisteredClaims{},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			t, err := token.SignedString([]byte("secret"))
			if err != nil {
				return err
			}
			return c.JSON(http.StatusOK, map[string]string{
				"token": t,
			})
		}
		return echo.ErrUnauthorized
	}
}

func InitUserRouter(e *echo.Echo) {
	e.POST("/login", Login())
}
