package main

import (
	"log"

	"github.com/FukeKazki/ta-dashboard/src/config"
	"github.com/FukeKazki/ta-dashboard/src/infrastracture"
	"github.com/FukeKazki/ta-dashboard/src/interface/handler"
	"github.com/FukeKazki/ta-dashboard/src/interface/router"
	"github.com/FukeKazki/ta-dashboard/src/usecase"
	"github.com/labstack/echo/v4"
)

func main() {
	log.Println("Starting TaDashboard API hello")
	e := echo.New()

	db, err := database.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	courseRepository := infrastracture.NewCourseRepository(db.Conn)
	courseUsecase := usecase.NewCourseUsecase(courseRepository)
	courseHandler := handler.NewcourseHandler(courseUsecase)
	router.InitCourseRouter(e, courseHandler)

	timeAttackRepository := infrastracture.NewTimeAttackRepository(db.Conn)
	timeAttackUsecase := usecase.NewTimeAttackUsecase(timeAttackRepository)
	timeAttackHandler := handler.NewTimeAttackHandler(timeAttackUsecase)
	router.InitTimeAttackRouter(e, timeAttackHandler)

	e.Logger.Fatal(e.Start(":4000"))
}
