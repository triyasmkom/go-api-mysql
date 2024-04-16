package helper

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Result struct {
	Rows    []map[string]interface{}
}

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:12345@tcp(localhost:3307)/test_db")
	if err != nil {
		panic(err)
		return nil
	}
	db.SetConnMaxIdleTime(time.Minute * 3)
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

func Query(query string, args ...interface{}) (string, error){
	var result []map[string]interface{}
	db := GetConnection()

	// Prepare the SQL statement with placeholders for parameters
	stmt, err := db.Prepare(query)
	if err != nil {
		return "", err
	}

	defer stmt.Close()

	// Execute the prepared statement with the parameter values
	rows, err := stmt.Query(args...)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	// Get column names// Get column names
	columns, err := rows.Columns()
	if err != nil{
		return "", err
	}

	// Iterate through the result set and store rows
	for rows.Next() {
		// Create a slice to hold the column values for this row
		values := make([]interface{}, len(columns))

		// Create a slice to hold pointers to each column value
		pointer := make([]interface{}, len(columns))
		for i := range values{
			pointer[i] = &values[i]
		}

		// Scan the row into the pointers slice
		if err := rows.Scan(pointer...); err != nil {
			return "", err
		}

		// Create a map to store column names and values for this row
		rowData := make(map[string]interface{})
		for i, column := range columns{
			rowData[column] = values[i]
		}
		// Append the row to the result
		result = append(result, rowData)
	}


	if err := rows.Err(); err != nil {
		return "", err
	}

	// Convert result struct to JSON
	jsonData, err := json.Marshal(result)
	if err != nil{
		return "", err
	}

	jsonStr := string(jsonData)
	// Unmarshal the JSON data into a slice of maps
	var data []map[string]string
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return "", err
	}

	// Decode base64-encoded strings
	for _, m := range data {
		for key, value := range m {
			decodedValue, err := base64.StdEncoding.DecodeString(value)
			if err != nil {
				fmt.Println("Error decoding base64:", err)
				return "", err
			}
			m[key] = string(decodedValue)
		}
	}

	// Marshal the updated data back to JSON
	updatedJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return "", err
	}

	return string(updatedJSON), nil
}
