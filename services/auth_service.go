package services

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	mysql "github.com/triyasmkom/rest-api-echo/helper"
	util "github.com/triyasmkom/rest-api-echo/helper"
	jwt "github.com/triyasmkom/rest-api-echo/middleware"
	database "github.com/triyasmkom/rest-api-echo/models/database"
	"github.com/triyasmkom/rest-api-echo/models/request"
	"github.com/triyasmkom/rest-api-echo/models/response"
)

func Register(context echo.Context) response.Body {
	body := new(request.Register)
	err := context.Bind(body)
	if err != nil {
		return response.Body{
			Status: false,
			Error: "Data kosong",
		}
	}
	data := database.User{
		Email: body.Email,
		Password: util.HashPassword(body.Password),
		Firstname: body.Firstname,
		Lastname: body.Lastname,
	}

	// Generate Token
	token, err := jwt.GenerateJwt(data)
	if err != nil {
		fmt.Println(err)
		return response.Body{
			Status: false,
			Error: "Error Generate Jwt",
		}
	}

	// Save User
	_, err = mysql.Query(
		"INSERT INTO auth_user (email, first_name, last_name, password) VALUES (?, ?, ?, ?)",
		data.Email,
		data.Firstname,
		data.Lastname,
		data.Password,
	)
	if err != nil {
		fmt.Println(err)
		return response.Body{
			Status: false,
			Error: "Register Failed",
		}
	}

	return response.Body{
		Status: true,
		Message: "Register Successfully",
		Data: token,
	}
}



func Login(context echo.Context) response.Body {
	body := new(request.Login)
	err := context.Bind(body)
	if err != nil {
		return response.Body{
			Status: false,
			Error: "Data konsong",
		}
	}

	// get user by email
	getUserByEmail, err := mysql.Query("SELECT first_name, last_name, password, email FROM auth_user WHERE email = ?", body.Email)
	if err != nil {
		if util.Debug() {
			fmt.Println(err)
		}

		return response.Body{
			Status: false,
			Error:  "Wrong Email or Password",
		}
	}

	fmt.Println("getUserByEmail: ", getUserByEmail)
	var user []database.User
	if err:= json.Unmarshal([]byte(getUserByEmail), &user); err != nil{
		if util.Debug() {
			fmt.Println(err)
		}

		return response.Body{
			Status: false,
			Error:  "Wrong Email or Password",
		}
	}
	if len(user) == 0 {
		if util.Debug() {
			fmt.Println("User tidak ada")
		}

		return response.Body{
			Status: false,
			Error:  "Wrong Email or Password",
		}
	}

	// Validasi password
	if _, err = util.VerifyPassword(body.Password, user[0].Password); err != nil {
		if util.Debug() {
			fmt.Println(err)
		}

		return response.Body{
			Status: false,
			Error:  "Wrong Email or Password",
		}
	}

	// Generate Jwt
	token, errToken := jwt.GenerateJwt(user[0])
	if errToken != nil {

		if util.Debug() {
			fmt.Println(errToken)
		}

		return response.Body{
			Status: false,
			Error:  "Error Generate Jwt",
		}
	}

	return response.Body{
		Status:  true,
		Message: "Login User Success",
		Data:    token,
	}
}
