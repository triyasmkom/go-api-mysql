package helper

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/triyasmkom/rest-api-echo/models/response"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
)

type PropertyDB struct {
	Username string
	Password string
	Host string
	Port string
	Name string
}
func GetPort() string {
	port := os.Getenv("PORT")

	if port == ""{
		return ":8080"
	}

	return ":" + port
}

func GetPropertyDB() PropertyDB {
	return PropertyDB{
		Username: os.Getenv("USERNAME_DB"),
		Password: os.Getenv("PASSWORD_DB"),
		Host: os.Getenv("HOST_DB"),
		Port: os.Getenv("PORT_DB"),
		Name: os.Getenv("NAME_DB"),
	}
}

func LoadEnv(path ...string)  {
	err := godotenv.Load(path...)
	if err != nil {
		fmt.Println("Load Env: ", err)
		panic("Failed load env file")
	}
}

func HashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func VerifyPassword(password string, hashPassword string) (response.Body, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return response.Body{}, err
	}
	return response.Body{}, nil
}

func Debug() bool {
	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		fmt.Println("Error Debug Variable")
		return false
	}
	return debug
}

func ParseJSONToMapSlice(jsonData string) ([]map[string]interface{}, error) {
	var data []map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
