package models

import "time"

// 데이터베이스 레이어

type Product struct {
	Image       string  `json:"img"`
	ImageAlt    string  `json:"img_alt"`
	Price       float64 `json:"price"`
	Promotion   float64 `json:"promotion"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
}

type Customer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	LoggedIn  bool   `json:"loggedin"`
}

type Order struct {
	Product
	Customer
	CustomerID   int       `json:"customer_id"`
	ProductID    int       `json:"product_id"`
	Price        float64   `json:"price"`
	PurchaseDate time.Time `json:"purchase_date"`
}
