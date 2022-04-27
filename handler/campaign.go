package handler

import (
	"campaign/campaign"
	"campaign/helper"
	"campaign/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	campagins, err := h.service.GetCampaigns(userId)
	if err != nil {
		response := helper.APIResponse("Error", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success", http.StatusOK, "success", campaign.FormatCampaigns(campagins))
	c.JSON(http.StatusOK, response)
}

// PAKAI PARAM/URI
func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput
	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helper.APIResponse("Error", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignData, err := h.service.GetCampaign(input)
	if err != nil {
		response := helper.APIResponse("Error", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success", http.StatusOK, "success", campaign.FormatCampaignDetail(campaignData))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := helper.FormatValidationError(err)
		response := helper.APIResponse("Gagal membuat campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// dapat dari JWT Middleware
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil {
		response := helper.APIResponse("Gagal membuat campaign", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Berhasil membuat campaign", http.StatusOK, "success", campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	var inputID campaign.GetCampaignDetailInput
	err := c.ShouldBindUri(&inputID)

	if err != nil {
		response := helper.APIResponse("Error", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input campaign.CreateCampaignInput
	err = c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := helper.FormatValidationError(err)
		response := helper.APIResponse("Gagal update campaign", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// dapat dari JWT Middleware
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	updatedCampaign, err := h.service.UpdateCampaign(inputID, input)
	if err != nil {
		response := helper.APIResponse("Gagal update campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Berhasil update campaign", http.StatusOK, "success", campaign.FormatCampaign(updatedCampaign))
	c.JSON(http.StatusOK, response)
}
