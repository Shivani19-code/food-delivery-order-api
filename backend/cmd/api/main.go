package main

import (
	"food-delivery-api/internal/handler"
	"food-delivery-api/internal/middleware"
	"food-delivery-api/internal/service"
	"food-delivery-api/pkg/config"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Initialize Database (including auto-seed)
	config.InitDB()

	// Initialize Services
	userService := &service.UserService{}
	restaurantService := &service.RestaurantService{}
	orderService := &service.OrderService{}

	// Initialize Handlers
	authHandler := handler.NewAuthHandler(userService)
	restaurantHandler := handler.NewRestaurantHandler(restaurantService)
	orderHandler := handler.NewOrderHandler(orderService)

	// Set up Router
	r := gin.Default()

	// Health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Public routes
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
	r.GET("/restaurants", restaurantHandler.GetAll)
	r.GET("/restaurants/:id", restaurantHandler.GetByID)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// Order routes
		protected.POST("/orders", orderHandler.PlaceOrder)
		protected.GET("/orders/:id", orderHandler.GetOrder)
		protected.GET("/orders/history", orderHandler.GetHistory)

		// Admin routes
		admin := protected.Group("/admin")
		admin.Use(middleware.AdminOnly())
		{
			admin.POST("/restaurants", restaurantHandler.Create)
			admin.POST("/restaurants/:id/menu", restaurantHandler.AddMenuItem)
			admin.PATCH("/orders/:id/status", orderHandler.UpdateStatus)
		}
	}

	// Start server
	log.Println("Server starting on :8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
