package router

import (
	"github.com/FukeKazki/ta-dashboard/src/interface/handler"
	"github.com/labstack/echo/v4"
)

func InitUserRouter(e *echo.Echo, userHandler handler.UserHandler) {
	e.POST("/login", userHandler.Login())
	e.POST("/signup", userHandler.Signup())
}
