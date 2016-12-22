package main

import (
	"github.com/KerwinKoo/gowauth/wdmethods"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.GET("/wifidog/ping/", wdmethods.Ping)
	e.GET("/wifidog/login/", wdmethods.Login)

	e.Logger.Fatal(e.Start(":8082"))
}
