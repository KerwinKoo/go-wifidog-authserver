package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func wdPing(c echo.Context) error {
	res := "pong"
	log.Println("answer:", res)
	return c.String(http.StatusOK, "pong")
}

func main() {
	e := echo.New()
	e.GET("/wifidog/ping/", wdPing)

	e.Logger.Fatal(e.Start(":8082"))
}
