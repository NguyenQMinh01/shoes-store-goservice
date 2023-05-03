package models

import "gorm.io/gorm"

// Product represents a product in the database
type Product struct {
    gorm.Model
    Name        string
    Description string
    Price       float64
    ImageURL    string
    // CategoryID  uint
    // Category    Category
}

type Order struct {
    gorm.Model
    UserID     uint        `json:"userId"`
    TotalPrice float32     `json:"totalPrice"`
    Items      []OrderItem `json:"items"`
}

type OrderItem struct {
    gorm.Model
    OrderID    uint    `json:"orderId"`
    ProductID  uint    `json:"productId"`
    Product    Product `json:"-"`
    Count      uint    `json:"count"`
    UnitPrice  float32 `json:"unitPrice"`
    TotalPrice float32 `json:"totalPrice"`
}

type User struct {
    gorm.Model
    Email    string `json:"email"`
    Password string `json:"-"`
}
