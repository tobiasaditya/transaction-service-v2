package main

import (
	"fmt"
	"log"
	"transaction-service-v2/handler"
	"transaction-service-v2/transaction"
	"transaction-service-v2/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "tobiasaditya:!D3papepe@tcp(127.0.0.1:3306)/personal_transaction?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connection to database success")

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	transactionRepo := transaction.NewRepository(db)
	transactionService := transaction.NewService(transactionRepo)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	router := gin.Default()

	api := router.Group("/api/v2")
	api.POST("/user/register", userHandler.RegisterUser)

	api.POST("/transaction/add", transactionHandler.CreateTransaction)
	api.GET("/transaction", transactionHandler.GetTransactionsUser)

	router.Run()
}
