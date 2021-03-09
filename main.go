package main

import (
	"go-gin-api/models"
	"go-gin-api/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	utils.PrepareDB(db)
	utils.GetFoodListOfRestaurant(db)
	
	// Read
	// var food models.Food
	// db.First(&food, 1)
	// db.First(&food, "name = ?", "Orange Cake")

	var restaurant models.Restaurant
	db.Find(&restaurant)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/restaurant", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"all": restaurant,
		})
	})
	r.GET("/restaurant/:resId/foods", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"resId": c.Param("resId"),
			"all" : food
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
