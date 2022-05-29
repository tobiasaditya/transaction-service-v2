package handler

import (
	"net/http"
	"transaction-service-v2/helper"
	"transaction-service-v2/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService: userService}
}

func (h userHandler) RegisterUser(c *gin.Context) {
	var input user.InputUser
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ErrorValidationResponse(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Invalid Input", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.CreateUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err}

		response := helper.APIResponse("Failed to register user", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
	}

	response := helper.APIResponse("Account has been registered", http.StatusOK, newUser)
	c.JSON(http.StatusOK, response)

}
