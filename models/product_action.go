package models

import (
	"demo-go-basic-backend/database"
)

func GetProducts() ([]Product, error) {
	// connect to DB
	db := database.InitDB()
	defer db.Close()

	// Query expect return rows
	rows, err := db.Query("SELECT id, name, description, quantity FROM products_tab")
	// check if query is error
	if err != nil {
		return nil, err
	}

	// create empty slices of Product
	products := []Product{}

	// do for every row
	for rows.Next() {
		var p Product

		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Quantity); err != nil {
			return nil, err
		}

		// append to products slice
		products = append(products, p)
	}

	return products, nil
}

func (p *Product) GetSingleProduct() error {
	//connect to DB
	db := database.InitDB()
	defer db.Close()

	return db.QueryRow("SELECT id, name, description, quantity FROM products_tab WHERE id = $1", p.ID).Scan(&p.ID, &p.Name, &p.Description, &p.Quantity)
}

func (p *Product) CreateProduct() error {
	// connect to DB
	db := database.InitDB()
	defer db.Close()

	return db.QueryRow("INSERT INTO products_tab (name, description, quantity) VALUES ($1, $2, $3) RETURNING id", p.Name, p.Description, p.Quantity).Scan(&p.ID)
}

func (p *Product) DeleteProduct() error {
	//connect to DB
	db := database.InitDB()
	defer db.Close()

	return db.QueryRow("DELETE FROM products_tab WHERE id = $1 RETURNING id", p.ID).Scan(&p.ID)
}

func (p *Product) UpdateProduct() error {
	//connect to DB
	db := database.InitDB()
	defer db.Close()

	return db.QueryRow("UPDATE products_tab SET name = $1, description = $2, quantity = $3 WHERE id = $4 RETURNING id", p.Name, p.Description, p.Quantity, p.ID).Scan(&p.ID)
}
