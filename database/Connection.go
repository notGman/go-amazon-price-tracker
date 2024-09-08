package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	DB_USER, exists := os.LookupEnv("DB_USER")
	if !exists {
		fmt.Println("DB_USER not found")
	}
	DB_PASS, exists := os.LookupEnv("DB_PASS")
	if !exists {
		fmt.Println("DB_PASS not found")
	}
	DB_NAME, exists := os.LookupEnv("DB_NAME")
	if !exists {
		fmt.Println("DB_NAME not found")
	}
	db, err := sql.Open("mysql", DB_USER+":"+DB_PASS+"@tcp(127.0.0.1:3306)/"+DB_NAME+"?parseTime=true")
	if err != nil {
		fmt.Println("Error:", err)
	}

	createProductTable := `
	CREATE TABLE IF NOT EXISTS products (
		id VARCHAR(255) PRIMARY KEY,
		name VARCHAR(255),
		url VARCHAR(255),
		email VARCHAR(255)
	);`

	createPriceTable := `
	CREATE TABLE IF NOT EXISTS prices (
		id VARCHAR(255),
		create_at TIMESTAMP,
		price INT
	);`

	_, err = db.Exec(createProductTable)
	if err != nil {
		log.Fatal("Failed to create table: ", err)
	}
	_, err = db.Exec(createPriceTable)
	if err != nil {
		log.Fatal("Failed to create table: ", err)
	}

	return db
}
