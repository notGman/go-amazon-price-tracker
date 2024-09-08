package scraper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/notgman/go-price/models"
)

const url = "https://amzn.in/d/"


func scraper(url string, product *models.ProductScrape) {

	c := colly.NewCollector(
		colly.AllowedDomains("www.amazon.in", "amzn.in"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnHTML("#productTitle", func(e *colly.HTMLElement) {
		if e.Index == 0 {
			product.Name = strings.TrimSpace(e.Text)
		}
	})

	c.OnHTML(".a-price-whole", func(e *colly.HTMLElement) {
		if e.Index == 0 {
			price := strings.ReplaceAll(e.Text, ".", "")
			product.Price, _ = strconv.Atoi(strings.ReplaceAll(price, ",", ""))
		}
	})

	err := c.Visit(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func ScrapeProduct() models.ProductScrape {
	product := models.ProductScrape{}

	fmt.Print("Enter the product ID: ")
	var ID string
	fmt.Scan(&ID)

	product.ID = ID

	scraper(url+ID, &product)

	fmt.Println("Product id: ", product.ID)
	fmt.Println("Product name: ", product.Name)
	fmt.Println("Product price: ", product.Price)

	return product
}

func ScrapeAllExistingProducts(products []string) []models.ProductScrape {
	result := []models.ProductScrape{}
	for _, productID := range products {
		product := models.ProductScrape{ID: productID}
		scraper(url+productID, &product)

		fmt.Println("Product id: ", product.ID)
		fmt.Println("Product name: ", product.Name)
		fmt.Println("Product price: ", product.Price)

		result = append(result, product)
	}

	return result
}
