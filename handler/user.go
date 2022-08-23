package handler

import (
	"fmt"
	"net/http"
	"transaction-service-v2/auth"
	"transaction-service-v2/email"
	"transaction-service-v2/helper"
	"transaction-service-v2/otp"
	"transaction-service-v2/user"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
	otpService  otp.Service
}

func NewUserHandler(userService user.Service, authService auth.Service, otpService otp.Service) *userHandler {
	return &userHandler{userService: userService, authService: authService, otpService: otpService}
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

	//Validate duplicate user
	foundUser, _ := h.userService.FindUserByPhone(input.PhoneNumber)
	if foundUser.ID != primitive.NilObjectID {
		errorMessage := gin.H{"errors": "duplicate username"}

		response := helper.APIResponse("Failed to register user", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//Generate otp
	newOtp, err := h.otpService.CreateOtp(input.Email)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed to register user", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//Send otp email
	err = email.SendEmailOtp(newOtp)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed to register user", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	format := otp.FormatOtp(newOtp)
	response := helper.APIResponse("Otp has been sent to "+newOtp.Receiver, http.StatusOK, format)
	c.JSON(http.StatusOK, response)
}

func (h userHandler) VerifyUser(c *gin.Context) {
	otpId := c.Param("otpId")
	var input user.InputVerifyUser
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ErrorValidationResponse(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Invalid Input", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//Verify Otp
	_, err = h.otpService.VerifyOtp(otpId, input.OtpValue)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed to verify new user", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	inputUser := user.InputUser{
		FullName:    input.FullName,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Password:    input.Password,
	}

	newUser, err := h.userService.CreateUser(inputUser)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed to register user", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	format := user.FormatUser(newUser)
	response := helper.APIResponse("Account has been registered", http.StatusOK, format)
	c.JSON(http.StatusOK, response)
}

func (h userHandler) LoginUser(c *gin.Context) {
	var input user.InputLogin
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ErrorValidationResponse(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Invalid Input", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	foundUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		fmt.Println(errorMessage)
		response := helper.APIResponse("Failed to login", http.StatusUnauthorized, errorMessage)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	//Generate token
	token, err := h.authService.GenerateToken(foundUser.ID.Hex())
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to generate token", http.StatusUnauthorized, errorMessage)
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	format := user.FormatLogin(token)
	response := helper.APIResponse("Success login", http.StatusOK, format)
	c.JSON(http.StatusOK, response)

}
