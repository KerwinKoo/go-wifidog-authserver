package wdmethods

import (
	"log"
	"net/http"

	"fmt"

	"github.com/KerwinKoo/gowauth/utils"
	"github.com/labstack/echo"
	"github.com/pquerna/ffjson/ffjson"
)

// LoginCtx wifidog login Context struct
type LoginCtx struct {
	GwAddress string `json:"gw_address"`
	GwID      string `json:"gw_id"`
	GwPort    string `json:"gw_port"`
	URL       string `json:"url"`
}

// Account user account
type Account struct {
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
}

// NewLoginCtx creator
func NewLoginCtx() *LoginCtx {
	lc := LoginCtx{}

	return &lc
}

// Ping response wifidog's Ping Get method in v1
func (loginCtx *LoginCtx) Ping(c echo.Context) error {
	res := "pong"
	log.Println("answer:", res)
	return c.String(http.StatusOK, "pong")
}

// Login wifidog login GET method
func (loginCtx *LoginCtx) Login(c echo.Context) error {
	loginCtx.GwAddress = c.QueryParam("gw_address")
	loginCtx.GwID = c.QueryParam("gw_id")
	loginCtx.GwPort = c.QueryParam("gw_port")
	loginCtx.URL = c.QueryParam("url")

	ctxJSON, err := ffjson.Marshal(loginCtx)
	if err != nil {
		panic(err)
	}
	log.Println("wifidog login:", string(ctxJSON))

	return c.Render(http.StatusOK, "wdargs", string(ctxJSON))
}

// LoginCheck wifi user login account check
func (loginCtx *LoginCtx) LoginCheck(c echo.Context) error {
	account := Account{}
	account.Name = c.QueryParam("name")
	account.Passwd = c.QueryParam("passwd")

	accountJSON, err := ffjson.Marshal(&account)
	if err != nil {
		panic(err)
	}

	log.Println("wifi account:", string(accountJSON))

	token := account.Name + account.Passwd
	gwAuthURL := fmt.Sprintf("http://%s:%s/auth?token=%s",
		loginCtx.GwAddress, loginCtx.GwPort, token)
	body, err := utils.HTTPGet(gwAuthURL)
	if err != nil {
		panic(err)
	}

	log.Println("router wifidog response body:", string(body))

	return c.String(http.StatusOK, string(body))
}

// Auth auth method
func (loginCtx *LoginCtx) Auth(c echo.Context) error {
	return c.String(http.StatusOK, "Auth: 1")
}

// Portal portal method
func (loginCtx *LoginCtx) Portal(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
}
