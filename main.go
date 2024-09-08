package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/notgman/go-price/functions"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Please provide a command (add, get, or update)")
		return
	}

	command := args[0]

	switch command {
	case "add":
		functions.AddProduct()
	case "get":
		functions.GetAllProducts()
	case "update":
		functions.UpdateProducts()
	default:
		fmt.Println("Invalid command. Please provide a valid command (add, get, or update)")
	}
}
