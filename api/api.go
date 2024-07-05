package api

import (
	"api-gateway-service/api/handler"
	"api-gateway-service/api/middleware"

	"github.com/gin-gonic/gin"
)

func Router(hand *handler.Handler) *gin.Engine {
    router := gin.Default()

    auth := router.Group("/api/auth")
    auth.POST("/register", hand.Register)
    auth.POST("/login", hand.Login)

    protected := router.Group("/")
	protected.Use(middleware.Middleware)

    user := protected.Group("/api/users")
    {
        user.GET("/:user-id", hand.GetUserByIdHandler)
        user.POST("/:user-id", hand.CreateUserHandler) 
        user.DELETE("/:user-id", hand.DeleteUserHandler)
        user.GET("/:user-id/profile", hand.GetUserByIdProfileHandler)
        user.PUT("/:user-id/profile", hand.UpdateUserProfileHandler)
    }



    garden := protected.Group("/api/gardens")
    {
        garden.POST("/", hand.CheckUser(), hand.CreateGardenHandler)
        garden.GET("/:garden-id", hand.ViewGardeHandler)
        garden.PUT("/:garden-id", hand.UpdateGardenHandler)
        garden.DELETE("/:garden-id", hand.DeleteGardenHandler)
        garden.POST("/:garden-id/plants", hand.AddPlanttoGarden)
        garden.GET("/:garden-id/plants", hand.ViewGardenPlantsHandler)
    }


    protected.GET("/api/users/:user-id/gardens", hand.ViewUserGardensHandler)


    

    plant := protected.Group("/api/plants")
    {
        plant.PUT("/:plant-id", hand.UpdatePlantHandler)
        plant.DELETE("/:plant-id", hand.DeletePlantHandler)
        plant.POST("/:plant-id/care-logs", hand.AddPlantCareLogHandler) 
        plant.GET("/:plant-id/care-logs", hand.ViewPlantCareLogsHandler) 
    }



    return router
}
