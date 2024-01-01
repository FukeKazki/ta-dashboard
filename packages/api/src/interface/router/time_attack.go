package router

import (
	"github.com/FukeKazki/ta-dashboard/src/interface/handler"
	"github.com/labstack/echo/v4"
)

func InitTimeAttackRouter(e *echo.Echo, timeAttackHandler handler.TimeAttackHandler) {
	e.GET("/dashboard", timeAttackHandler.FindUserTARecord())
}
