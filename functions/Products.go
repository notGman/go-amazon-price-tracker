package functions

import (
	"fmt"

	"github.com/notgman/go-price/database"
	"github.com/notgman/go-price/mail"
	"github.com/notgman/go-price/scraper"
)

func AddProduct() {
	product := scraper.ScrapeProduct()
	fmt.Print("Enter your email: ")
	var email string
	fmt.Scan(&email)

	database.AddProduct(product, email)
}

func GetAllProducts() {
	products := database.GetAllProducts()
	for _, product := range products {
		fmt.Println(product)
	}

	prices := database.GetAllPrices()
	for _, price := range prices {
		fmt.Println(price)
	}
}

func UpdateProducts() {
	productIds := database.GetAllProductID()
	scrapedData := scraper.ScrapeAllExistingProducts(productIds)

	for _, product := range scrapedData {
		price := database.GetLatestPrice(product.ID)
		if price > product.Price {
			mail.SendMail(product, product.Price)
		}
	}

	database.UpdateTimeStamp(scrapedData)
}
