package handler

import (
	"food-delivery-api/internal/model"
	"food-delivery-api/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	orderService *service.OrderService
}

func NewOrderHandler(os *service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: os}
}

func (h *OrderHandler) PlaceOrder(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from middleware
	userID, _ := c.Get("userID")
	order.UserID = userID.(uint)

	if err := h.orderService.PlaceOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID, _ := c.Get("userID")

	order, err := h.orderService.GetOrder(uint(id), userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) GetHistory(c *gin.Context) {
	userID, _ := c.Get("userID")
	orders, err := h.orderService.GetUserOrders(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) UpdateStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.orderService.UpdateStatus(uint(id), req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Status updated"})
}
