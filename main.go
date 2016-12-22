package main

import (
	"github.com/KerwinKoo/gowauth/wdmethods"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/wifidog/ping/", wdmethods.Ping)
	e.GET("/wifidog/login/", wdmethods.Login)

	e.Logger.Fatal(e.Start(":8082"))
}
