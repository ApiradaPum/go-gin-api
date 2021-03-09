package models

import "gorm.io/gorm"

type FoodOfRestaurant struct {
	gorm.Model
	ResId  uint
	FoodId uint
}
