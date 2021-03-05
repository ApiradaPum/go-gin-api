package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	Name  string
	Price uint
	Amount uint
}

type Restaurant struct {
	gorm.Model
	Name  string
	LocationId uint
}

type FoodOfRestaurant struct {
	gorm.Model
	ResId  uint
	FoodId uint
}

type Location struct {
	gorm.Model
	LocationAddress string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Food{})
	db.AutoMigrate(&Restaurant{})
	db.AutoMigrate(&Location{})

	// Create
	db.Create(&Food{Name: "Orange Cake", Price: 100, Amount:200})
	db.Create(&Food{Name: "Chocolate Cake", Price: 120, Amount:50})
	db.Create(&Food{Name: "Strawberry Cake", Price: 120, Amount:100})
	db.Create(&Food{Name: "Brownie", Price: 140, Amount:250})

	db.Create(&Restaurant{Name: "Mint Sweet", LocationId: 1})
	db.Create(&Restaurant{Name: "Chocolate House", LocationId: 2})

	db.Create(&FoodOfRestaurant{ResId: 1, FoodId: 1})
	db.Create(&FoodOfRestaurant{ResId: 1, FoodId: 2})
	db.Create(&FoodOfRestaurant{ResId: 2, FoodId: 2})
	db.Create(&FoodOfRestaurant{ResId: 1, FoodId: 3})
	db.Create(&FoodOfRestaurant{ResId: 2, FoodId: 4})

	db.Create(&Location{LocationAddress: "MBK Bangkok"})
	db.Create(&Location{LocationAddress: "101 True Digital Park"})
	// Read
	var food Food
	db.First(&food, 1)                 
	db.First(&food, "name = ?", "Orange Cake") 

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/welcome", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"food": food,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}