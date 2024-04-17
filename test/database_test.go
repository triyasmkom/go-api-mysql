package test

import (
	"database/sql"
	"fmt"
	database "github.com/triyasmkom/rest-api-echo/helper"
	model "github.com/triyasmkom/rest-api-echo/models/database"
	"testing"
)

func TestGetDB(t *testing.T) {
	database.LoadEnv("./../.env.dev")
	result, err := database.Query("SELECT first_name, last_name FROM auth_user WHERE id = 1", )
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}

func TestInsertDB(t *testing.T) {
	database.LoadEnv("./../.env.dev")
	_, err := database.Query("INSERT INTO auth_user(email, first_name, last_name) VALUES ( ?, ?, ?);", "test5@gmail.com", "test2", "test2")

	if err != nil {
		panic(err)
	}
}

func TestUpdateDB(t *testing.T) {
	database.LoadEnv("./../.env.dev")
	_, err := database.Query("UPDATE auth_user SET email = ? WHERE id = ? ", "test6@gmail.com", 1)

	if err != nil {
		panic(err)
	}
}

func TestDeleteDB(t *testing.T) {
	database.LoadEnv("./../.env.dev")
	_, err := database.Query("DELETE FROM auth_user WHERE id = ? ", 7)

	if err != nil {
		panic(err)
	}
}

func TestDB(t *testing.T)  {
	var result []model.User
	database.LoadEnv("./../.env.dev")
	propertyDB := database.GetPropertyDB()
	if err := database.Connect(propertyDB); err != nil{
		fmt.Println("Error Connection to mysql: ", err)
		return
	}

	defer database.CloseConnection()

	rows, err := database.DB.Query("SELECT first_name, last_name, email FROM auth_user")
	if err != nil {
		fmt.Println("Error executing query: ", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var user model.User
		// Scan data dari baris ke variabel
		err := rows.Scan(&user.Firstname, &user.Lastname, &user.Email)
		if err != nil {
			// Handle error
			fmt.Println("Error scanning row:", err)
			return
		}

		result = append(result, user)
	}

	fmt.Println(result)
}

func TestTransactionDB(t *testing.T) {
	database.LoadEnv("./../.env.dev")
	propertyDB := database.GetPropertyDB()
	if err := database.Connect(propertyDB); err != nil{
		fmt.Println("Error Connection to mysql: ", err)
		return
	}

	defer database.CloseConnection()

	err := database.ExecuteTransaction(database.DB, func(tx *sql.Tx) error {
		_, err := tx.Exec("INSERT INTO auth_user(email, first_name, last_name) VALUES ( ?, ?, ?);", "test5@gmail.com", "test2", "test2")
		if err !=nil{
			return err
		}
		var result [] model.User
		rows, err := tx.Query("SELECT id, first_name, last_name, email FROM auth_user WHERE email = ? ", "test5@gmail.com")
		if err !=nil{
			return err
		}

		for rows.Next() {
			var user model.User
			// Scan data dari baris ke variabel
			err := rows.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Email)
			if err != nil {
				// Handle error
				fmt.Println("Error scanning row:", err)
				return err
			}

			result = append(result, user)
		}

		_, err = tx.Exec("UPDATE auth_user SET email = ? WHERE id = ? ", "test10@gmail.com", result[0].Id)
		if err !=nil{
			return err
		}

		_, err = tx.Exec("DELETE FROM auth_user WHERE id = ? ", result[0].Id)
		if err !=nil{
			return err
		}
		return nil
	})

	if err !=nil {
		fmt.Println("Error Transaction")
		return
	}

	fmt.Println("Transaction executed successfully")
}

