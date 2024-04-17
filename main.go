package main

import (
	"github.com/labstack/echo/v4"
	util "github.com/triyasmkom/rest-api-echo/helper"
	route "github.com/triyasmkom/rest-api-echo/routes"
	"os"
	"strconv"
)

func main()  {
	var envProd bool
	util.LoadEnv()
	getEnv, err := strconv.ParseBool(os.Getenv("ENV_PROD"))
	if err != nil {
		envProd = false
	} else {
		envProd = getEnv
	}

	if envProd {
		util.LoadEnv("./.env.prod")
	}

	util.LoadEnv("./.env.dev")

	app := echo.New()
	route.Init(app)
	app.Logger.Fatal(app.Start(util.GetPort()))
}