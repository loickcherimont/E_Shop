package main

import (
	"fmt"
	"net/http"
)

// *** FOR DATABASE ***
func getAllProducts() ([]Product, error) {
	// An products slice to hold data from returned rows.
	var products []Product

	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var prod Product
		if err := rows.Scan(&prod.ID, &prod.Title, &prod.Img, &prod.Price, &prod.Quantity); err != nil {
			return nil, fmt.Errorf("error: %v", err)
		}
		products = append(products, prod)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	return products, nil
}

// *** FOR TEMPLATES ***
func indexHandler(w http.ResponseWriter, r *http.Request) {
	data, err := getAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}
