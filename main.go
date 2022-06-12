package main

import (
	"fmt"
	"net/http"
	"strings"
	"transaction-service-v2/auth"
	"transaction-service-v2/config"
	"transaction-service-v2/handler"
	"transaction-service-v2/helper"
	"transaction-service-v2/transaction"
	"transaction-service-v2/user"

	"github.com/dgrijalva/jwt-go"
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
	trxCollection := newDatabase.GetCollection("trx_collection")
	fmt.Println("Connection to database success")

	authService := auth.NewJwtService()

	userRepo := user.NewRepository(userCollection)
	userService := user.NewService(userRepo)
	userHandler := handler.NewUserHandler(userService, authService)

	transactionRepo := transaction.NewRepository(trxCollection)
	transactionService := transaction.NewService(transactionRepo)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	router := gin.Default()

	api := router.Group("/api/v2")
	api.POST("/user/register", userHandler.RegisterUser)
	api.POST("/user/login", userHandler.LoginUser)

	api.POST("/transaction/add", authMiddleware(authService, userService), transactionHandler.CreateTransaction)
	api.GET("/transaction", authMiddleware(authService, userService), transactionHandler.GetTransactionsUser)

	router.Run(":8000")
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil header, Authorization Bearer tokentokentoken
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			errorMessage := gin.H{"errors": "Session expired"}
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, errorMessage)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		var tokenString string
		headerToken := strings.Split(authHeader, " ")

		if len(headerToken) == 2 {
			tokenString = headerToken[1]
		}

		// Validasi token
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			errorMessage := gin.H{"errors": "Session expired"}
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, errorMessage)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//Kalau valid, decode token, buat dapetin claim/payload dari jwt token
		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			errorMessage := gin.H{"errors": "Session expired"}
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, errorMessage)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := claim["user_id"].(string)

		//Get user by Id
		foundUser, err := userService.FindUserByID(userId)
		if err != nil {
			errorMessage := gin.H{"errors": "Session expired"}
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, errorMessage)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		fmt.Println(foundUser.ID)
		c.Set("currentUser", foundUser)
	}

}
