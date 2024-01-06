package router

import (
	"github.com/FukeKazki/ta-dashboard/src/interface/handler"
	"github.com/labstack/echo/v4"
)

func InitTimeAttackRouter(e *echo.Echo, timeAttackHandler handler.TimeAttackHandler) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Name
	e.GET("/dashboard", timeAttackHandler.FindUserTARecord())
}
