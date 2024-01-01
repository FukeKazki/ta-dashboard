package router

import (
	"github.com/FukeKazki/ta-dashboard/src/interface/handler"
	"github.com/labstack/echo/v4"
)

func InitCourseRouter(e *echo.Echo, courseHandler handler.CourseHandler) {
	e.GET("/course", courseHandler.GetAll())
}
