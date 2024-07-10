package main

import (
	"log"

	"github.com/faridlan/firestore-vercel/config"
	"github.com/faridlan/firestore-vercel/controller"
	"github.com/faridlan/firestore-vercel/helper"
	"github.com/faridlan/firestore-vercel/repository"
	"github.com/faridlan/firestore-vercel/service"
)

func main() {

	err := helper.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	client, err := config.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo, client)
	userController := controller.NewUserController(userService)

	router := config.Router{
		UserController: userController,
	}

	app := config.NewRouter(router)

	app.Listen(":3000")

}
