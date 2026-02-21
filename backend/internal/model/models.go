package model

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `json:"name" binding:"required"`
	Email     string         `gorm:"uniqueIndex" json:"email" binding:"required,email"`
	Password  string         `json:"password,omitempty" binding:"required"`
	Role      string         `json:"role" gorm:"default:'customer'"` // admin or customer
}

type Restaurant struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Name        string         `json:"name" binding:"required"`
	Description string         `json:"description"`
	Address     string         `json:"address" binding:"required"`
	Rating      float32        `json:"rating" gorm:"default:0"`
	MenuItems   []MenuItem     `json:"menu_items,omitempty"`
}

type MenuItem struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	RestaurantID uint           `json:"restaurant_id"`
	Name         string         `json:"name" binding:"required"`
	Description  string         `json:"description"`
	Price        float64        `json:"price" binding:"required"`
}

type Order struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	UserID       uint           `json:"user_id"`
	RestaurantID uint           `json:"restaurant_id"`
	TotalPrice   float64        `json:"total_price"`
	Status       string         `json:"status" gorm:"default:'pending'"` // pending, confirmed, preparing, delivered
	OrderItems   []OrderItem    `json:"order_items"`
}

type OrderItem struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	OrderID    uint    `json:"order_id"`
	MenuItemID uint    `json:"menu_item_id"`
	Quantity   int     `json:"quantity" binding:"required"`
	Price      float64 `json:"price"` // Fixed price at time of order
}
