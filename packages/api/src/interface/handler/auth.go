package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/FukeKazki/ta-dashboard/src/config"
	"github.com/FukeKazki/ta-dashboard/src/domain/model"
	"github.com/FukeKazki/ta-dashboard/src/usecase"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthHandler interface {
	Signup() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type authHandler struct {
	userUsecase usecase.UserUsecase
}

func NewAuthHandler(userUsecase usecase.UserUsecase) AuthHandler {
	return &authHandler{userUsecase: userUsecase}
}

func (h *authHandler) Signup() echo.HandlerFunc {
	return func(c echo.Context) error {
		userName := c.FormValue("username")
		password := c.FormValue("password")

		user, err := h.userUsecase.CreateUser(userName, password)
		// 詰め変える
		userDto := &model.UserDto{
			Name: user.Name.Value(),
		}
		// create jwt token
		fmt.Println(user.Id.String())
		claims := &config.JwtCustomClaims{
			user.Name.Value(),
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte("secret"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"token": t,
			"user":  userDto,
		})
	}
}

func (h *authHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		userName := c.FormValue("username")
		password := c.FormValue("password")
		user, err := h.userUsecase.FindByName(userName)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		// 詰め変える
		userDto := &model.UserDto{
			Name: user.Name.Value(),
		}

		if user.Password == password {
			claims := &config.JwtCustomClaims{
				user.Name.Value(),
				jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			t, err := token.SignedString([]byte("secret"))
			if err != nil {
				return err
			}
			return c.JSON(http.StatusOK, map[string]interface{}{
				"token": t,
				"user":  userDto,
			})
		}
		return echo.ErrUnauthorized
	}
}
