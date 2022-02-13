package handler

import (
	"campaign/helper"
	"campaign/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := helper.FormatValidationError(err)
		response := helper.APIResponse("User gagal registrasi", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("User gagal registrasi", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := user.FormatUser(newUser, "token")

	response := helper.APIResponse("User berhasil registrasi", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := helper.FormatValidationError(err)
		response := helper.APIResponse("User gagal login", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err, httpResponse := h.userService.Login(input)
	if err != nil {
		response := helper.APIResponse("User gagal login", httpResponse, "error", err.Error())
		c.JSON(httpResponse, response)
		return
	}

	data := user.FormatUser(newUser, "token")

	response := helper.APIResponse("User berhasil login", httpResponse, "success", data)
	c.JSON(httpResponse, response)
}
