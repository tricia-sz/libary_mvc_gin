package main

import (
	bookcontroller "libarymvc/internal/books/controllers"
	usercontroller "libarymvc/internal/users/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	bookController := bookcontroller.NewBooksController()
	userController := usercontroller.NewUserController()

	bookController.RegisterRouter(router)
	userController.RegisterRouter(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)

	}
}
