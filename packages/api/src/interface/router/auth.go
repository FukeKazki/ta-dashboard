package router

import (
	"github.com/FukeKazki/ta-dashboard/src/interface/handler"
	"github.com/labstack/echo/v4"
)

func InitAuthRouter(e *echo.Echo, authHandler handler.AuthHandler) {
	e.POST("/login", authHandler.Login())
	e.POST("/signup", authHandler.Signup())
}
