package main

import (
	"fmt"
	"transaction-service-v2/config"
	"transaction-service-v2/handler"
	"transaction-service-v2/user"

	"github.com/gin-gonic/gin"
)

func main() {
	// dsn := "tobiasaditya:!D3papepe@tcp(127.0.0.1:3306)/personal_transaction?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	newDatabase := config.NewDatabase()
	// Check the connection
	userCollection := newDatabase.GetCollection("user_collection")
	// trxCollection := newDatabase.GetCollection("trx_collection")
	fmt.Println("Connection to database success")

	userRepo := user.NewRepository(userCollection)
	userService := user.NewService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// transactionRepo := transaction.NewRepository(trxCollection)
	// transactionService := transaction.NewService(transactionRepo)
	// transactionHandler := handler.NewTransactionHandler(transactionService)

	router := gin.Default()

	api := router.Group("/api/v2")
	api.POST("/user/register", userHandler.RegisterUser)
	api.POST("/user/login", userHandler.LoginUser)

	// api.POST("/transaction/add", transactionHandler.CreateTransaction)
	// api.GET("/transaction", transactionHandler.GetTransactionsUser)

	router.Run(":8000")
}
