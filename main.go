package main

import (
    "log"
    "net/http"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"

    "github.com/gin-gonic/gin"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Order struct {
    gorm.Model
    Email string
    Phone string
}

type Product struct {
    ID    int
    Name  string
    Price float32
}

func main() {
    db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=minhlatoi28092001 dbname=rails_react_webbangiay_test port=5432 sslmode=disable"), &gorm.Config{})
    if err != nil {
        log.Fatalf("Can't connect to database: %v", err)
    }
    

    // Migrate the schema
    db.AutoMigrate(&Order{})
    db.AutoMigrate(&Product{})

    r := gin.Default()

    // Create a new order
    r.POST("/orders", func(c *gin.Context) {
        var order Order
        c.BindJSON(&order)

        // Save the order to the database
        result := db.Create(&order)
        if result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Order created"})
    })

    r.Run(":8080")
}
