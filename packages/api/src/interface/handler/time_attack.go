package handler

import (
	"net/http"

	"github.com/FukeKazki/ta-dashboard/src/usecase"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

type TimeAttackHandler interface {
	FindUserTARecord() echo.HandlerFunc
}

type timeAttackHandler struct {
	timeAttackUsecase usecase.TimeAttackUsecase
}

func NewTimeAttackHandler(timeAttackUsecase usecase.TimeAttackUsecase) TimeAttackHandler {
	return &timeAttackHandler{timeAttackUsecase: timeAttackUsecase}
}

func (h *timeAttackHandler) FindUserTARecord() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtCustomClaims)
		name := claims.Name

		foundedcourses, err := h.timeAttackUsecase.FindUserTARecord()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, foundedcourses)
	}
}
