package handler

import (
	"net/http"

	"github.com/FukeKazki/ta-dashboard/src/usecase"
	"github.com/labstack/echo/v4"
)

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
		foundedcourses, err := h.timeAttackUsecase.FindUserTARecord()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, foundedcourses)
	}
}
