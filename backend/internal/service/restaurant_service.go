package service

import (
	"food-delivery-api/internal/model"
	"food-delivery-api/pkg/config"
	"gorm.io/gorm"
)

type RestaurantService struct{}

func (s *RestaurantService) GetAll() ([]model.Restaurant, error) {
	var restaurants []model.Restaurant
	err := config.DB.Find(&restaurants).Error
	return restaurants, err
}

func (s *RestaurantService) GetByID(id uint) (*model.Restaurant, error) {
	var restaurant model.Restaurant
	err := config.DB.Preload("MenuItems").First(&restaurant, id).Error
	if err != nil {
		return nil, err
	}
	return &restaurant, nil
}

func (s *RestaurantService) Create(restaurant *model.Restaurant) error {
	return config.DB.Create(restaurant).Error
}

func (s *RestaurantService) Update(restaurant *model.Restaurant) error {
	return config.DB.Save(restaurant).Error
}

func (s *RestaurantService) Delete(id uint) error {
	return config.DB.Delete(&model.Restaurant{}, id).Error
}

// Menu Items
func (s *RestaurantService) AddMenuItem(item *model.MenuItem) error {
	return config.DB.Create(item).Error
}

func (s *RestaurantService) GetMenu(restaurantID uint) ([]model.MenuItem, error) {
	var items []model.MenuItem
	err := config.DB.Where("restaurant_id = ?", restaurantID).Find(&items).Error
	return items, err
}
