package handler

import (
	"api-gateway-service/generated/gardenManagement"
	"api-gateway-service/generated/users"
)


type Handler struct{
	User users.UserManagementClient
	Garden gardenManagement.GardenManagementClient
}


func NewHendler(user users.UserManagementClient, garden gardenManagement.GardenManagementClient) *Handler {
	return &Handler{
		User: user,
		Garden: garden,
	}
}