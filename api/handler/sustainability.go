package handler

import (
	"api-gateway-service/generated/sustainability"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) LogImpactHandle(ctx *gin.Context) {
	var logImpact sustainability.LogImpactRequest

	err := ctx.ShouldBindJSON(&logImpact)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	resp, err := h.Sustainability.LogImpact(ctx, &logImpact)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"Success": resp.Success,
		"Message": "Log impact created successfully",
	})
}

func (h *Handler) GetUserImpactHandle(ctx *gin.Context) {
	userId := ctx.Param("user-id")

	lomImpact, err := h.Sustainability.GetUserImpact(ctx, &sustainability.GetUserImpactRequest{UserId: userId})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, lomImpact)
}

func (h *Handler) GetCommunityImpactHandle(ctx *gin.Context) {
	id := ctx.Param("community-id")

	logImpact, err := h.Sustainability.GetCommunityImpact(ctx, &sustainability.GetCommunityImpactRequest{Id: id})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, logImpact)

}

func (h *Handler) GetChallengesHandle(ctx *gin.Context) {
	challenges, err := h.Sustainability.GetChallenges(ctx, &sustainability.GetChallengesRequest{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, challenges)
}

func (h *Handler) JoinChallengeHandler(ctx *gin.Context) {

}
