package wdmethods

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// Ping response wifidog's Ping Get method in v1
func Ping(c echo.Context) error {
	res := "pong"
	log.Println("answer:", res)
	return c.String(http.StatusOK, "pong")
}
