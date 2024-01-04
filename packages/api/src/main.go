package main

import (
	"log"

	"github.com/FukeKazki/ta-dashboard/src/config"
	"github.com/FukeKazki/ta-dashboard/src/infrastracture"
	"github.com/FukeKazki/ta-dashboard/src/interface/handler"
	"github.com/FukeKazki/ta-dashboard/src/interface/router"
	"github.com/FukeKazki/ta-dashboard/src/usecase"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

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

	r := e
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	r.Use(echojwt.WithConfig(config))

	timeAttackRepository := infrastracture.NewTimeAttackRepository(db.Conn)
	timeAttackUsecase := usecase.NewTimeAttackUsecase(timeAttackRepository)
	timeAttackHandler := handler.NewTimeAttackHandler(timeAttackUsecase)
	router.InitTimeAttackRouter(r, timeAttackHandler)

	e.Logger.Fatal(e.Start(":4000"))
}
