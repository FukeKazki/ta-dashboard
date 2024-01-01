package handler

import (
	"net/http"

	"github.com/FukeKazki/ta-dashboard/src/usecase"
	"github.com/labstack/echo/v4"
)

type CourseHandler interface {
	GetAll() echo.HandlerFunc
}

type courseHandler struct {
	courseUsecase usecase.CourseUsecase
}

func NewcourseHandler(courseUsecase usecase.CourseUsecase) CourseHandler {
	return &courseHandler{courseUsecase: courseUsecase}
}

func (h *courseHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		foundedcourses, err := h.courseUsecase.FindAll()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, foundedcourses)
	}
}
