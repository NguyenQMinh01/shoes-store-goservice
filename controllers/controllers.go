package controllers

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "/home/minhnguyen/webbangiay/goservice/models"
)

// ProductsController is the controller for products
type ProductsController struct {
    db *gorm.DB
}

// Index shows a list of products
func (pc ProductsController) Index(c *gin.Context) {
    var products []models.Product
    pc.db.Find(&products)

    c.JSON(http.StatusOK, gin.H{
        "products": products,
    })
}
