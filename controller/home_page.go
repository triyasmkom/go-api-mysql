package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HomePage(context echo.Context) error {
	return context.String(http.StatusOK, "HELLO WORLD")
}
