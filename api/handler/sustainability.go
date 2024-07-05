package handler

import (

	pb "api-gateway-service/generated/sustainability"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) LogImpactHandle(ctx *gin.Context) {
	var logImpact pb.LogImpactRequest

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

	lomImpact, err := h.Sustainability.GetUserImpact(ctx, &pb.GetUserImpactRequest{UserId: userId})

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

	logImpact, err := h.Sustainability.GetCommunityImpact(ctx, &pb.GetCommunityImpactRequest{Id: id})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, logImpact)

}

func (h *Handler) GetChallengesHandle(ctx *gin.Context) {
	challenges, err := h.Sustainability.GetChallenges(ctx, &pb.GetChallengesRequest{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, challenges)
}

func (h *Handler) JoinChallengeHendler1(ctx *gin.Context) {
	var req pb.JoinChallengeRequest

	err := ctx.BindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	resp, err := h.Sustainability.JoinChallenge(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateChallengeProgressHendler1(ctx *gin.Context) {
	var req pb.UpdateChallengeProgressRequest

	err := ctx.BindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	resp, err := h.Sustainability.UpdateChallengeProgress(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetUserChallengesHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.Sustainability.GetUserChallenges(ctx, &pb.GetUserChallengesRequest{UserId: id})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetUsersLeaderboardHandler(ctx *gin.Context) {
	resp, err := h.Sustainability.GetCommunitiesLeaderboard(ctx, &pb.GetCommunitiesLeaderboardRequest{})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetCommunitiesLeaderboardHandler(ctx *gin.Context) {
	var req pb.GetCommunitiesLeaderboardRequest
	err := ctx.BindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	resp, err := h.Sustainability.GetCommunitiesLeaderboard(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
