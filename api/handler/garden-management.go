package handler

import (
	"api-gateway-service/generated/gardenManagement"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateGardenHandler(ctx *gin.Context) {
	var garden gardenManagement.CreateGardenRequest

	err := ctx.ShouldBindJSON(&garden)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})

		return 
	}

	resp, err := h.Garden.CreateGarden(ctx, &garden)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

func (h *Handler) ViewGardeHandler(ctx *gin.Context) {
	id := ctx.Param("garden-id")
	req := gardenManagement.ViewGardenRequest{Id: id}

	garden, err := h.Garden.ViewGarden(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, garden)
}

func (h *Handler) UpdateGardenHandler(ctx *gin.Context) {
	id := ctx.Param("garden-id")

	var updateGarden gardenManagement.UpdateGardenRequest
	
	err := ctx.ShouldBindJSON(&updateGarden)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})

		return 
	}

	updateGarden.Id = id

	resp, err := h.Garden.UpdateGarden(ctx, &updateGarden)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	if !resp.Success {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": resp.Success,
			"message": "Garden not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": resp.Success,
		"message": "Garden updeted successfully",
	})
}