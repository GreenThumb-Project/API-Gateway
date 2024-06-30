package api

import (
	"api-gateway-service/api/handler"

	"github.com/gin-gonic/gin"
)

func Router(hand *handler.Handler) *gin.Engine {
	router := gin.Default()

	return router
}