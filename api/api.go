package api

import (
	"api-gateway-service/api/handler"

	"github.com/gin-gonic/gin"
)

func Router(hand *handler.Handler) *gin.Engine {
	router := gin.Default()

	user := router.Group("api/user")
	{
		user.GET("/{id}", hand.GetUserByIdHandler)
		user.PUT("/{id}", hand.CreateUserHandler)
		user.DELETE("/{id}", hand.DeleteUserHandler)
		user.GET("/{id}/prpfile", hand.GetUserByIdProfileHandler)
		user.PUT("/{id}/profile", hand.UpdateUserProfileHandler)

	}

	auth:=router.Group("/api/auth/")
	{
		auth.POST("registr",hand.CreateUserHandler)
	}

	return router
}
