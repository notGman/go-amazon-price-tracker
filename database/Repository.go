package database

import (
	"fmt"

	"github.com/notgman/go-price/models"
)

func AddProduct(product models.ProductScrape, Email string) {
	URL := "https://amzn.in/d/" + product.ID
	insert, err := AddProcuctToDB(product.ID, product.Name, URL, Email)
	if err != nil {
		fmt.Println("Error ", err)
	}
	defer insert.Close()

	timeStamp, err := InsertTimeStamp(product.ID, product.Price)
	if err != nil {
		fmt.Println("Error ", err)
	}
	defer timeStamp.Close()
}

func GetAllProducts() []models.Product {
	products := []models.Product{}

	rows, err := GetAllProductsFromDB()
	if err != nil {
		fmt.Println("Error ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.URL, &product.Email)
		if err != nil {
			fmt.Println("Error ", err)
		}
		products = append(products, product)
	}
	return products
}

func GetAllPrices() []models.Price {
	prices := []models.Price{}

	rows, err := GetAllPricesFromDB()
	if err != nil {
		fmt.Println("Error ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var price models.Price
		err := rows.Scan(&price.ID, &price.CreatedAt, &price.Price)
		if err != nil {
			fmt.Println("Error ", err)
		}
		prices = append(prices, price)
	}
	return prices
}

func GetAllProductID() []string {
	data := []string{}
	rows, err := GetAllProductIDFromDB()
	if err != nil {
		fmt.Println("Error ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var ID string
		err := rows.Scan(&ID)
		if err != nil {
			fmt.Println("Error ", err)
		}
		data = append(data, ID)
	}
	return data
}

func UpdateTimeStamp(scrapedData []models.ProductScrape) {
	for _, product := range scrapedData {

		insert, err := InsertTimeStamp(product.ID, product.Price)
		if err != nil {
			fmt.Println("Error ", err)
		}
		defer insert.Close()
	}
	fmt.Println("Timestamps updated")
}

func GetUserMail(ID string) string {
	row, err := GetUserMailFromDB(ID)
	if err != nil {
		fmt.Println("Error ", err)
	}
	defer row.Close()

	var email string
	for row.Next() {
		err := row.Scan(&email)
		if err != nil {
			fmt.Println("Error ", err)
		}
	}
	return email
}

func GetLatestPrice(ID string) int {
	row, err := GetLatestPriceFromDB(ID)
	if err != nil {
		fmt.Println("Error ", err)
	}
	defer row.Close()

	var price int
	for row.Next() {
		err := row.Scan(&price)
		if err != nil {
			fmt.Println("Error ", err)
		}
	}
	return price
}
