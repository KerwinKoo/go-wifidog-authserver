package wdmethods

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pquerna/ffjson/ffjson"
)

// LoginCtx wifidog login Context struct
type LoginCtx struct {
	GwAddress string `json:"gw_address"`
	GwID      string `json:"gw_id"`
	GwPort    string `json:"gw_port"`
}

// Ping response wifidog's Ping Get method in v1
func Ping(c echo.Context) error {
	res := "pong"
	log.Println("answer:", res)
	return c.String(http.StatusOK, "pong")
}

// Login wifidog login GET method
func Login(c echo.Context) error {
	loginCtx := LoginCtx{}
	loginCtx.GwAddress = c.QueryParam("gw_address")
	loginCtx.GwID = c.QueryParam("gw_id")
	loginCtx.GwPort = c.QueryParam("gw_port")

	ctxJson, err := ffjson.Marshal(&loginCtx)
	if err != nil {
		panic(err)
	}
	log.Println("wifidog login:", string(ctxJson))
	return c.String(http.StatusOK, string(ctxJson))
}
