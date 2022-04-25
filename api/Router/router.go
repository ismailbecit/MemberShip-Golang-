package Router

import (
	"app/api/config"
	"app/api/controller"

	"github.com/labstack/echo/v4"
)

func Router() {
	config.Conn()
	e := echo.New()
	e.POST("/login", controller.LoginUser)
	e.POST("/register", controller.RegisterUser)
	e.DELETE("/user", controller.DelUser)

	e.Start(":8080")

}
