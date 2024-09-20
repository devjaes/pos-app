// cmd/api/main.go
package main

import (
	"log"

	"github.com/devjaes/pos-app/internal/customer"
	"github.com/devjaes/pos-app/internal/product"
	"github.com/devjaes/pos-app/internal/sale"
	"github.com/devjaes/pos-app/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize Gin router
	r := gin.Default()

	// Initialize services and handlers
	productService := product.NewService(db)
	productHandler := product.NewHandler(productService)
	product.RegisterRoutes(r, productHandler)

	customerService := customer.NewService(db)
	customerHandler := customer.NewHandler(customerService)
	customer.RegisterRoutes(r, customerHandler)

	saleService := sale.NewService(db)
	saleHandler := sale.NewHandler(saleService)
	sale.RegisterRoutes(r, saleHandler)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
