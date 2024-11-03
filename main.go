package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
	"warehouse-management/config"
	"warehouse-management/controllers"
	"warehouse-management/repositories"
	"warehouse-management/routes"
	"warehouse-management/services"
)

func main() {
	db := config.ConnectDatabase()
	warehouseRepo := repositories.NewWarehouseRepository(db)
	warehouseService := services.NewWarehouseService(warehouseRepo)
	warehouseController := controllers.NewWarehouseController(warehouseService)

	router := gin.Default()
	routes.RegisterRoutes(router, warehouseController)

	// Configure CORS middleware
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(config))

	router.Run(":8080")
}
