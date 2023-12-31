package router

import (
	"github.com/FukeKazki/ta-dashboard/src/interface/handler"
	"github.com/labstack/echo/v4"
)

func InitStageRouter(e *echo.Echo, stageHandler handler.StageHandler) {
	e.GET("/stages", stageHandler.GetAll())
}
