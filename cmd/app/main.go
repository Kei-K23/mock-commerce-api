package main

import "github.com/Kei-K23/go-ecommerce-api/db"

func main() {
	// Start the db connection
	db.ConnectDB()

	// Close the db connection when before main function out of scope
	defer db.Pool.Close()
}
