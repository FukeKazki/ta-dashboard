package main

import (
	"log"

	"github.com/FukeKazki/ta-dashboard/src/interface/handler"
	"github.com/FukeKazki/ta-dashboard/src/interface/router"
	"github.com/labstack/echo/v4"
)

func main() {
	log.Println("Starting TaDashboard API hello")
	e := echo.New()

	stageHandler := handler.NewStageHandler()
	router.InitStageRouter(e, stageHandler)

	e.Logger.Fatal(e.Start(":4000"))
}
