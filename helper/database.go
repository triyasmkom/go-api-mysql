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

var DB *sql.DB

func GetConnection() *sql.DB {
	propertyDB := GetPropertyDB()
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", propertyDB.Username, propertyDB.Password, propertyDB.Host, propertyDB.Port, propertyDB.Name)

	db, err := sql.Open("mysql", dataSourceName)
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

func Connect(propertyDB PropertyDB) error {
	// Format string untuk koneksi
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", propertyDB.Username, propertyDB.Password, propertyDB.Host, propertyDB.Port, propertyDB.Name)

	// Membuat koneksi database
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}

	// Memeriksa koneksi database
	err = db.Ping()
	if err != nil {
		return err
	}

	// Menetapkan koneksi sebagai variabel global
	DB = db
	fmt.Println("Connected to MySQL database!")
	return nil
}


// CloseConnection digunakan untuk menutup koneksi database
func CloseConnection() {
	if DB != nil {
		DB.Close()
	}
	fmt.Println("Connection to MySQL database closed.")
}


// executeTransaction adalah fungsi bantu untuk melakukan transaksi dengan database MySQL.
func ExecuteTransaction(db *sql.DB, txFunc func(*sql.Tx) error) (err error) {
	// Mulai transaksi
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Defer fungsi Rollback untuk menjalankannya jika terjadi kesalahan atau panic.
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Panic kembali setelah rollback
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit() // Commit transaksi jika tidak ada kesalahan
		}
	}()

	// Panggil fungsi txFunc dengan transaksi yang disediakan sebagai argumen.
	err = txFunc(tx)
	return err
}


