package test

import (
	"fmt"
	database "github.com/triyasmkom/rest-api-echo/helper"
	"testing"
)

func TestGetDB(t *testing.T) {
	result, err := database.Query("SELECT first_name, last_name FROM auth_user WHERE id = 1", )
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}

func TestInsertDB(t *testing.T) {
	_, err := database.Query("INSERT INTO auth_user(email, first_name, last_name) VALUES ( ?, ?, ?);", "test5@gmail.com", "test2", "test2")

	if err != nil {
		panic(err)
	}
}

func TestUpdateDB(t *testing.T) {
	_, err := database.Query("UPDATE auth_user SET email = ? WHERE id = ? ", "test6@gmail.com", 1)

	if err != nil {
		panic(err)
	}
}

func TestDeleteDB(t *testing.T) {
	_, err := database.Query("DELETE FROM auth_user WHERE id = ? ", 7)

	if err != nil {
		panic(err)
	}
}


