package main

import (
	"github.com/labstack/echo/v4"
	util "github.com/triyasmkom/rest-api-echo/helper"
	route "github.com/triyasmkom/rest-api-echo/routes"

)

func main()  {
	util.LoadEnv()
	app := echo.New()
	route.Init(app)
	app.Logger.Fatal(app.Start(util.GetPort()))
}