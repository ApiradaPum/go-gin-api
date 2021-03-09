package utils

import (
	"go-gin-api/models"

	"gorm.io/gorm"
)

func PrepareDB(db *gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(&models.Food{})
	db.AutoMigrate(&models.Restaurant{})
	db.AutoMigrate(&models.FoodOfRestaurant{})
	db.AutoMigrate(&models.Location{})

	// Create
	db.Create(&models.Food{Name: "Orange Cake", Price: 100, Amount: 200})
	db.Create(&models.Food{Name: "Chocolate Cake", Price: 120, Amount: 50})
	db.Create(&models.Food{Name: "Strawberry Cake", Price: 120, Amount: 100})
	db.Create(&models.Food{Name: "Brownie", Price: 140, Amount: 250})

	db.Create(&models.Restaurant{Name: "Mint Sweet", LocationId: 1})
	db.Create(&models.Restaurant{Name: "Chocolate House", LocationId: 2})

	db.Create(&models.FoodOfRestaurant{ResId: 1, FoodId: 1})
	db.Create(&models.FoodOfRestaurant{ResId: 1, FoodId: 2})
	db.Create(&models.FoodOfRestaurant{ResId: 2, FoodId: 2})
	db.Create(&models.FoodOfRestaurant{ResId: 1, FoodId: 3})
	db.Create(&models.FoodOfRestaurant{ResId: 2, FoodId: 4})

	db.Create(&models.Location{LocationAddress: "MBK Bangkok"})
	db.Create(&models.Location{LocationAddress: "101 True Digital Park"})
}

func GetFoodListOfRestaurant(db *gorm.DB, resId Int) []*models.Food {
	var food []*models.Food
	var foodOfRestaurant []*models.FoodOfRestaurant
	queryResult := db.Select("id").Find(&foodOfRestaurant, "resId = ?", resId)
	db.Find(&food, "id = ?", foodOfRestaurant)
	queryResult2 := db.Find(&food, listFood)
	
	return 
}
