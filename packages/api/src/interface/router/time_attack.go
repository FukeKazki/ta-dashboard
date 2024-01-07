package router

import (
	"github.com/FukeKazki/ta-dashboard/src/interface/handler"
	"github.com/labstack/echo/v4"
)

func InitTimeAttackRouter(e *echo.Group, timeAttackHandler handler.TimeAttackHandler) {
	e.GET("/ta/:userName", timeAttackHandler.FindUserTARecord())
	e.POST("/ta/:userName/:id", timeAttackHandler.UpdateTARecord())
}
