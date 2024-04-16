package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/triyasmkom/rest-api-echo/services"
	"net/http"
)

func Register(context echo.Context) error {
	save := services.Register(context)
	if !save.Status {
		return context.JSON(http.StatusBadRequest, save)
	}
	return context.JSON(http.StatusCreated, save)
}


func Login(context echo.Context) error {
	save := services.Login(context)
	if !save.Status {
		return context.JSON(http.StatusBadRequest, save)
	}
	return context.JSON(http.StatusCreated, save)
}
