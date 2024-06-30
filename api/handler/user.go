package handler

import (
	pb "api-gateway-service/generated/users"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *Handler)  GetUserByIdHandler(c *gin.Context){
	userId:=c.Param("userId")

	resp,err:=h.User.GetUserById(c,&pb.GetUserByIdRequest{UserId: userId})

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK,resp)
	return
}