package config

import (
	"food-delivery-api/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("food_delivery.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto Migration
	err = DB.AutoMigrate(&model.User{}, &model.Restaurant{}, &model.MenuItem{}, &model.Order{}, &model.OrderItem{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	SeedData()
}

func SeedData() {
	var count int64
	DB.Model(&model.Restaurant{}).Count(&count)
	if count > 0 {
		return // Data already seeded
	}

	restaurants := []model.Restaurant{
		{
			Name:        "Burger King",
			Description: "Home of the Whopper",
			Address:     "123 Street, City",
			Rating:      4.5,
			MenuItems: []model.MenuItem{
				{Name: "Whopper", Description: "Classic burger", Price: 5.99},
				{Name: "Fries", Description: "Crispy golden fries", Price: 2.99},
			},
		},
		{
			Name:        "Pizza Hut",
			Description: "Best pizza in town",
			Address:     "456 Avenue, City",
			Rating:      4.2,
			MenuItems: []model.MenuItem{
				{Name: "Pepperoni Pizza", Description: "Classic pizza", Price: 12.99},
				{Name: "Garlic Bread", Description: "Bread with garlic", Price: 4.99},
			},
		},
	}

	for _, r := range restaurants {
		DB.Create(&r)
	}
	log.Println("Database seeded with sample restaurants and menu items.")
}
