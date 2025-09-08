package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swag/model"
)

type userResponse struct {
	Data []models.User `json:"data"`
}

type countryResponse struct {
	Countries []models.Country `json:"countries"`
}

type errResponse struct {
	Message string `json:"message"`
}

type create struct {
	Data models.User `json:"data"`
}

// GetUsers return list of all users from the database
// @Summary return list of all users from the database
// @Description list of all users
// @Tags Users
// @Success 200 {object} userResponse
// @Router /users [get]
func GetUsers(context *gin.Context) {
	users := models.ListUsers()

	context.JSON(http.StatusOK, userResponse{Data: users})

}

// CreateUsers creates new user
// @Summary creates new user
// @Description creates new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body  models.User  true "User"
// @Security bearerToken
// @Success 201 {object} models.User
// @Failure 400 {object} errResponse
// @Router /users [post]
func CreateUsers(context *gin.Context) {
	var req models.User
	err := context.ShouldBindJSON(&req)

	if err != nil {
		context.JSON(http.StatusBadRequest, errResponse{Message: err.Error()})
		return
	}

	err = models.CreateUsers(req)
	if err != nil {
		context.JSON(http.StatusInternalServerError, errResponse{Message: err.Error()})
		return
	}

	context.JSON(http.StatusCreated, create{Data: req})
}

// Get coutries of users
// @Summary Get countries of users
// @Description Get countries of users
// @Tags Country
// @Success 200 {object} countryResponse
// @Failure 500 {object} errResponse
// @Router /country [get]
func ListCountries(context *gin.Context) {
	countries, err := models.ListCountries()
	if err != nil {
		context.JSON(http.StatusInternalServerError, errResponse{Message: err.Error()})
		return
	}

	context.JSON(http.StatusOK, countryResponse{Countries: countries})

}
