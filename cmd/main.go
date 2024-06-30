package main

import (
	"api-gateway-service/api"
	"log"
)

func main() {
	hand := NewConnect()

	router := api.Router(hand)

	log.Fatal(router.Run(":8080"))
}
