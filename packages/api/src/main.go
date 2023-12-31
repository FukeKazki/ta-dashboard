package main

import (
	"log"

	"github.com/FukeKazki/ta-dashboard/src/interface/handler"
	"github.com/FukeKazki/ta-dashboard/src/interface/router"
	"github.com/FukeKazki/ta-dashboard/src/usecase"
	"github.com/labstack/echo/v4"
)

func main() {
	log.Println("Starting TaDashboard API hello")
	e := echo.New()

	stageUsecase := usecase.NewStageUsecase()
	stageHandler := handler.NewStageHandler(stageUsecase)
	router.InitStageRouter(e, stageHandler)

	e.Logger.Fatal(e.Start(":4000"))
}
