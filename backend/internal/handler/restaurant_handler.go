package handler

import (
	"food-delivery-api/internal/model"
	"food-delivery-api/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RestaurantHandler struct {
	restaurantService *service.RestaurantService
}

func NewRestaurantHandler(rs *service.RestaurantService) *RestaurantHandler {
	return &RestaurantHandler{restaurantService: rs}
}

func (h *RestaurantHandler) GetAll(c *gin.Context) {
	restaurants, err := h.restaurantService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, restaurants)
}

func (h *RestaurantHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	restaurant, err := h.restaurantService.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Restaurant not found"})
		return
	}
	c.JSON(http.StatusOK, restaurant)
}

func (h *RestaurantHandler) Create(c *gin.Context) {
	var restaurant model.Restaurant
	if err := c.ShouldBindJSON(&restaurant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.restaurantService.Create(&restaurant); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, restaurant)
}

func (h *RestaurantHandler) AddMenuItem(c *gin.Context) {
	restaurantID, _ := strconv.Atoi(c.Param("id"))
	var item model.MenuItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item.RestaurantID = uint(restaurantID)

	if err := h.restaurantService.AddMenuItem(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, item)
}
