package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type StageHandler interface {
	GetAll() echo.HandlerFunc
}

type stageHandler struct {
}

func NewStageHandler() StageHandler {
	return &stageHandler{}
}

func (h *stageHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}
}
