package main

import (
	"api-gateway-service/api"
	"api-gateway-service/api/handler"
	"api-gateway-service/config"
	"api-gateway-service/generated/gardenManagement"
	"api-gateway-service/generated/users"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg := config.Load()
	hand := NewConnect()

	router := api.Router(hand)

	log.Fatal(router.Run(cfg.HTTP_PORT))
}

func NewConnect() *handler.Handler {
	usersConn, err := grpc.NewClient("localhost:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Error: ", err.Error())
		return nil
	}

	usersService := users.NewUserManagementClient(usersConn)

	gardenManagementConn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("ERROR: ", err.Error())
		return nil
	}

	gardenManagementService := gardenManagement.NewGardenManagementClient(gardenManagementConn)

	return &handler.Handler{
		User:   usersService,
		Garden: gardenManagementService,
	}
}
