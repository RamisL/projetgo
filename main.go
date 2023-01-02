package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	/*dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "user:password@tcp(127.0.0.1:3306)/projetgo?charset=utf8mb4&parseTime=True&loc=Local"
	}

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&product.Product{})

	taskRepository := product.NewRepository(db)*/
	//taskService := product.NewService(taskRepository)
	//taskHandler := handler.NewTaskHandler(taskService)

	pl("Hello, World!")
	//router := gin.Default()
	//router.GET("/products")
	//router.run("localhost:3035")
}
