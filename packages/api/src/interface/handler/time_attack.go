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
	CreateTARecord() echo.HandlerFunc
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

func (h *timeAttackHandler) CreateTARecord() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (h *timeAttackHandler) UpdateTARecord() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: ここでjwtの認証を行う
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*config.JwtCustomClaims)
		userID := claims.UserId
		fmt.Println(user, claims, userID)
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		u := new(model.TimeAttack)
		if err := c.Bind(u); err != nil {
			return err
		}

		err = h.timeAttackUsecase.UpdateTARecord(id, u)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, "success")
	}
}
