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

	courseUsecase := usecase.NewCourseUsecase()
	courseHandler := handler.NewcourseHandler(courseUsecase)
	router.InitCourseRouter(e, courseHandler)

	e.Logger.Fatal(e.Start(":4000"))
}
