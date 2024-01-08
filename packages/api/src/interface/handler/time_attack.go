package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/FukeKazki/ta-dashboard/src/config"
	"github.com/FukeKazki/ta-dashboard/src/domain/model"
	"github.com/FukeKazki/ta-dashboard/src/usecase"
	"github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v4"
)

type TimeAttackHandler interface {
	FindUserTARecord() echo.HandlerFunc
	UpdateTARecord() echo.HandlerFunc
}

type timeAttackHandler struct {
	timeAttackUsecase usecase.TimeAttackUsecase
}

func NewTimeAttackHandler(timeAttackUsecase usecase.TimeAttackUsecase) TimeAttackHandler {
	return &timeAttackHandler{timeAttackUsecase: timeAttackUsecase}
}

func (h *timeAttackHandler) FindUserTARecord() echo.HandlerFunc {
	return func(c echo.Context) error {
		userName := c.Param("userName")

		foundedcourses, err := h.timeAttackUsecase.FindUserTARecord(userName)
		fmt.Println(foundedcourses)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, foundedcourses)
	}
}

func (h *timeAttackHandler) UpdateTARecord() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*config.JwtCustomClaims)
		userName := claims.UserName
		fmt.Println(user, claims, userName)
		paramUserName := c.Param("userName")
		if userName != paramUserName {
			return c.JSON(http.StatusBadRequest, "invalid user")
		}
		recordId, err := strconv.Atoi(c.Param("recordId"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		u := new(model.TimeAttack)
		if err := c.Bind(u); err != nil {
			return err
		}

		err = h.timeAttackUsecase.UpdateTARecord(recordId, u)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, "success")
	}
}
