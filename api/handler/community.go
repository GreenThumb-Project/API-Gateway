package handler

import (
	"api-gateway-service/generated/community"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCommunityHenndler(ctx *gin.Context) {
	var req *community.CreateCommunityRequest

	err := ctx.ShouldBind(&req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	resp, err := h.Community.CreateCommunity(ctx, req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, &resp)
}

func (h *Handler) GetCommunityHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	req := community.GetCommunityRequest{Id: id}

	resp, err := h.Community.GetCommunity(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)

}

func (h *Handler) UpdateCommunityHendler(ctx *gin.Context) {
	var req *community.UpdateCommunityRequest

	err := ctx.ShouldBind(&req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	resp, err := h.Community.UpdateCommunity(ctx, req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteCommunityHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	req := community.DeleteCommunityRequest{Id: id}

	resp, err := h.Community.DeleteCommunity(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) ListCommunitiesHandler(ctx *gin.Context) {
	resp, err := h.Community.ListCommunities(ctx, &community.ListCommunitiesRequest{})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) JoinCommunityHendler(ctx *gin.Context) {
	id :=ctx.Param("id")

	req:=community.JoinCommunityRequest{Id: id}

	resp,err:=h.Community.JoinCommunity(ctx,&req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,resp)
}

func (h *Handler) LeaveCommunityHandler(ctx *gin.Context)  {
	id :=ctx.Param("id")

	req:=community.LeaveCommunityRequest{Id: id}

	resp,err:=h.Community.LeaveCommunity(ctx,&req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,resp)

}

func (h *Handler) CreateCommunityEventHendler(ctx *gin.Context)  {
	var req *community.CreateCommunityEventRequest

	err:=ctx.BindJSON(req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	resp,err:=h.Community.CreateCommunityEvent(ctx,req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,resp)
}

func (h *Handler) ListCommunityEventsHandler(ctx *gin.Context)  {
	id :=ctx.Param("id")

	req:=community.ListCommunityEventsRequest{Id: id}

	resp,err:=h.Community.ListCommunityEvents(ctx,&req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,resp)

}


func (h *Handler) CreateCommunityForumPostHendler(ctx *gin.Context)  {
	var req *community.CreateCommunityForumPostRequest

	err:=ctx.BindJSON(req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	resp,err:=h.Community.CreateCommunityForumPost(ctx,req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,resp)
}

func (h *Handler) ListCommunityForumPostsHandler(ctx *gin.Context)  {
	id :=ctx.Param("id")

	req:=community.ListCommunityForumPostsRequest{ComunityId: id}

	resp,err:=h.Community.ListCommunityForumPosts(ctx,&req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,resp)
}

func (h *Handler) AddForumPostCommentHendler(ctx *gin.Context)  {
	var req *community.AddForumPostCommentRequest

	err:=ctx.BindJSON(req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	resp,err:=h.Community.AddForumPostComment(ctx,req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,resp)
}


func (h *Handler) ListForumPostCommentsHandler(ctx *gin.Context)  {
	id :=ctx.Param("id")

	req:=community.ListForumPostCommentsRequest{PostId: id}

	resp,err:=h.Community.ListForumPostComments(ctx,&req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,resp)
}