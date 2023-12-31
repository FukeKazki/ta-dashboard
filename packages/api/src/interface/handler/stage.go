package handler

import (
	"net/http"

	"github.com/FukeKazki/ta-dashboard/src/usecase"
	"github.com/labstack/echo/v4"
)

type StageHandler interface {
	GetAll() echo.HandlerFunc
}

type stageHandler struct {
	stageUsecase usecase.StageUsecase
}

func NewStageHandler(stageUsecase usecase.StageUsecase) StageHandler {
	return &stageHandler{stageUsecase: stageUsecase}
}

func (h *stageHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		foundedStages, err := h.stageUsecase.FindAll()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, foundedStages)
	}
}
