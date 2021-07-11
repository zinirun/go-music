package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// 데이터베이스 레이어

type Product struct {
	gorm.Model
	Image       string  `json:"img"`
	ImageAlt    string  `gorm:"img_alt" json:"img_alt"`
	Price       float64 `json:"price"`
	Promotion   float64 `gorm:"promotion" json:"promotion"`
	ProductName string  `gorm:"product_name" json:"product_name"`
	Description string
}

func (Product) TableName() string {
	return "products"
}

type Customer struct {
	gorm.Model
	Name      string  `json:"name"`
	FirstName string  `gorm:"column:first_name" json:"first_name"`
	LastName  string  `gorm:"column:last_name" json:"last_name"`
	Email     string  `gorm:"column:email" json:"email"`
	Password  string  `gorm:"column:password" json:"password"`
	LoggedIn  bool    `gorm:"column:loggedin" json:"loggedin"`
	Orders    []Order `json:"orders"`
}

func (Customer) TableName() string {
	return "customers"
}

type Order struct {
	gorm.Model
	Product
	Customer
	CustomerID   int       `gorm:"column:customer_id"`
	ProductID    int       `gorm:"column:product_id"`
	Price        float64   `gorm:"column:price" json:"price"`
	PurchaseDate time.Time `gorm:"column:purchase_date" json:"purchase_date"`
}

func (Order) TableName() string {
	return "orders"
}
