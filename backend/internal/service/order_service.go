package service

import (
	"errors"
	"food-delivery-api/internal/model"
	"food-delivery-api/pkg/config"
	"gorm.io/gorm"
)

type OrderService struct{}

func (s *OrderService) PlaceOrder(order *model.Order) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		var total float64 = 0

		for i := range order.OrderItems {
			var menuItem model.MenuItem
			if err := tx.First(&menuItem, order.OrderItems[i].MenuItemID).Error; err != nil {
				return errors.New("menu item not found")
			}
			// Set price at time of order
			order.OrderItems[i].Price = menuItem.Price
			total += menuItem.Price * float64(order.OrderItems[i].Quantity)
		}

		order.TotalPrice = total
		order.Status = "pending"

		if err := tx.Create(order).Error; err != nil {
			return err
		}

		return nil
	})
}

func (s *OrderService) GetOrder(id uint, userID uint) (*model.Order, error) {
	var order model.Order
	err := config.DB.Preload("OrderItems").Where("id = ? AND user_id = ?", id, userID).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (s *OrderService) GetUserOrders(userID uint) ([]model.Order, error) {
	var orders []model.Order
	err := config.DB.Preload("OrderItems").Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

func (s *OrderService) UpdateStatus(id uint, status string) error {
	return config.DB.Model(&model.Order{}).Where("id = ?", id).Update("status", status).Error
}
