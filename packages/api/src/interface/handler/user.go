package handler

import (
	"net/http"

	"github.com/FukeKazki/ta-dashboard/src/usecase"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	Signup() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}
type jwtCustomClaims struct {
	userId string `json:"userId"`
	jwt.RegisteredClaims
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: userUsecase}
}

func (h *userHandler) Signup() echo.HandlerFunc {
	return func(c echo.Context) error {
		userName := c.FormValue("username")
		password := c.FormValue("password")

		user, err := h.userUsecase.CreateUser(userName, password)
		// create jwt token
		claims := &jwtCustomClaims{
			user.Id.String(),
			jwt.RegisteredClaims{},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte("secret"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}
}

func (h *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		// userName := c.FormValue("username")
		// password := c.FormValue("password")

		// if userName == "jon" && password == "shhh!" {
		// 	claims := &jwtCustomClaims{
		// 		"Jon Snow",
		// 		true,
		// 		jwt.RegisteredClaims{},
		// 	}
		// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// 	t, err := token.SignedString([]byte("secret"))
		// 	if err != nil {
		// 		return err
		// 	}
		// 	return c.JSON(http.StatusOK, map[string]string{
		// 		"token": t,
		// 	})
		// }
		return echo.ErrUnauthorized
	}
}
