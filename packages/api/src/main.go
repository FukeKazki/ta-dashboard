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

func main() {
	log.Println("Starting TaDashboard API hello")
	e := echo.New()

	db, err := config.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	userRepository := infrastracture.NewUserRepository(db.Conn)
	timeAttackRepository := infrastracture.NewTimeAttackRepository(db.Conn)

	userUsecase := usecase.NewUserUsecase(userRepository, timeAttackRepository)
	authHandler := handler.NewAuthHandler(userUsecase)
	router.InitAuthRouter(e, authHandler)

	r := e.Group("/api")

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(config.JwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	r.Use(echojwt.WithConfig(config))

	timeAttackUsecase := usecase.NewTimeAttackUsecase(timeAttackRepository)
	timeAttackHandler := handler.NewTimeAttackHandler(timeAttackUsecase)
	router.InitTimeAttackRouter(r, timeAttackHandler)

	e.Logger.Fatal(e.Start(":4000"))
}
