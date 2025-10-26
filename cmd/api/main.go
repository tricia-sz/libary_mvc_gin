package main

import (
	"libarymvc/internal/users/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	usersController := controllers.NewUserController()
	usersController.RegisterRouter(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)

	}
}
