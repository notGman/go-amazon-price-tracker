package models

import "time"

type Product struct {
	ID   string
	Name string
	URL  string `json:"url"`
	Email string
}

type Price struct {
	ID        string
	CreatedAt time.Time `json:"created_at"`
	Price     int
}

type ProductScrape struct {
	ID    string
	Name  string
	Price int
}

type EmailData struct {
	Name string
	Price int
}
