package database

import "database/sql"

func AddProcuctToDB(ID string, Name string, URL string, Email string) (*sql.Rows, error) {
	db := ConnectDB()

	insert, err := db.Query("INSERT INTO products (id, name, url, email) VALUES (?, ?, ?, ?)", ID, Name, URL, Email)
	return insert, err
}

func GetAllProductsFromDB() (*sql.Rows, error) {
	db := ConnectDB()

	rows, err := db.Query("SELECT * FROM products")
	return rows, err
}

func GetAllPricesFromDB() (*sql.Rows, error) {
	db := ConnectDB()

	rows, err := db.Query("SELECT * FROM prices")
	return rows, err
}

func GetProductFromDB(ID string) (*sql.Rows, error) {
	db := ConnectDB()

	row, err := db.Query("SELECT * FROM products WHERE id = ?", ID)
	return row, err
}

func InsertTimeStamp(ID string, Price int) (*sql.Rows, error) {
	db := ConnectDB()

	insert, err := db.Query("INSERT INTO prices (id, price) VALUES (?, ?)", ID, Price)
	return insert, err
}

func GetAllProductIDFromDB() (*sql.Rows, error) {
	db := ConnectDB()

	rows, err := db.Query("SELECT id FROM products")
	return rows, err
}

func GetUserMailFromDB(ID string) (*sql.Rows, error) {
	db := ConnectDB()

	row, err := db.Query("SELECT email FROM products WHERE id = ?", ID)
	return row, err
}

func GetLatestPriceFromDB(ID string) (*sql.Rows, error) {
	db := ConnectDB()

	row, err := db.Query("SELECT price FROM prices WHERE id = ? ORDER BY created_at DESC LIMIT 1", ID)
	return row, err
}
