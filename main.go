package main

import (
	"html/template"
	"io"

	"github.com/KerwinKoo/gowauth/wdmethods"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Template init
type Template struct {
	templates *template.Template
}

// Render base render
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("template/*.html")),
	}

	e := echo.New()
	e.Renderer = t

	e.Static("/static", "static")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}\n",
	}))

	loginCtx := wdmethods.NewLoginCtx()
	wdGroup := e.Group("/wifidog")
	// wdGroup.Use(middleware.AddTrailingSlash())
	wdGroup.GET("/ping/", loginCtx.Ping)
	wdGroup.GET("/login/", loginCtx.Login)
	wdGroup.GET("/logincheck", loginCtx.LoginCheck)
	wdGroup.GET("/auth/", loginCtx.Auth)
	wdGroup.GET("/portal/", loginCtx.Portal)

	e.Logger.Fatal(e.Start(":8082"))
}
