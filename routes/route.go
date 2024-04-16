package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	auth "github.com/triyasmkom/rest-api-echo/controller"
	page "github.com/triyasmkom/rest-api-echo/controller"
)

func Init(app *echo.Echo)  {
	app.Use(middleware.Logger())
	app.GET("/", page.HomePage)
	app.POST("/auth/register", auth.Register)
	app.POST("/auth/login", auth.Login)
}
