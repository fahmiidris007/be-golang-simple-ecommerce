package handlers

import (
	"simple-ecommerce/models"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []models.Product
		var wg sync.WaitGroup

		//add 1 goroutine to WaitGroup
		wg.Add(1)

		//start goroutine
		go func() {
			defer wg.Done() //goroutine is done
			//get list of products from database
			db.Find(&products)
		}()

		//wait until all goroutines are done
		wg.Wait()

		c.JSON(200, products)
	}

}

func GetProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product models.Product
		if err := db.First(&product, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Product not found"})
			return
		}
		c.JSON(200, product)
	}
}

func CreateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.Product
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}
		db.Create(&input)
		c.JSON(201, input)
	}
}

func UpdateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product models.Product
		if err := db.First(&product, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Product not found"})
			return
		}
		var input models.Product
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}
		db.Model(&product).Updates(input)
		c.JSON(200, product)
	}
}

func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product models.Product
		if err := db.First(&product, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Product not found"})
			return
		}
		db.Delete(&product)
		c.JSON(200, gin.H{"message": "Product deleted"})
	}
}
