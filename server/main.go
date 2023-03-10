package main

import (
	"log"
	"os"

	"github.com/RamisL/server/adapter"
	"github.com/RamisL/server/broadcast"
	"github.com/RamisL/server/handler"
	"github.com/RamisL/server/payment"
	"github.com/RamisL/server/product"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//var pl = fmt.Println

func main() {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "user:password@tcp(127.0.0.1:3306)/projectgo?charset=utf8mb4&parseTime=True&loc=Local"
	}

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&product.Product{}, &payment.Payment{})

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	paymentRepository := payment.NewRepository(db)
	paymentService := payment.NewService(paymentRepository)
	broadcaster := broadcast.NewBroadcaster(20)
	paymentHandler := handler.NewPaymentHandler(paymentService, broadcaster)
	adapter := adapter.NewGinAdapter(broadcaster, paymentService)

	r := gin.Default()
	api := r.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api.GET("/stream", adapter.Stream)

	api.POST("/product/create", productHandler.CreateProduct)
	api.GET("/product/showAll", productHandler.GetAllProduct)
	api.GET("/product/show/:id", productHandler.GetByIdProduct)
	api.PUT("/product/update/:id", productHandler.UpdateProduct)
	api.DELETE("/product/delete/:id", productHandler.DeleteProduct)

	api.POST("/payment/create", paymentHandler.CreatePayment)
	api.GET("/payment/showAll", paymentHandler.GetAllPayment)
	api.GET("/payment/show/:id", paymentHandler.GetByIdPayment)
	api.PUT("/payment/update/:id", paymentHandler.UpdatePayment)
	api.DELETE("/payment/delete/:id", paymentHandler.DeletePayment)

	r.Run(":3000")
}
