package main

import (
	"github.com/KerwinKoo/gowauth/wdmethods"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/wifidog/ping/", wdmethods.WdPing)

	e.Logger.Fatal(e.Start(":8082"))
}
