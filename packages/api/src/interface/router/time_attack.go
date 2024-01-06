package router

import (
	"github.com/FukeKazki/ta-dashboard/src/interface/handler"
	"github.com/labstack/echo/v4"
)

func InitTimeAttackRouter(e *echo.Group, timeAttackHandler handler.TimeAttackHandler) {
	e.GET("/ta", timeAttackHandler.FindUserTARecord())
	e.POST("/ta/:id", timeAttackHandler.UpdateTARecord())
}
