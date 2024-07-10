package api

import (
	"log"
	"net/http"

	"github.com/faridlan/firestore-vercel/config"
	"github.com/faridlan/firestore-vercel/controller"
	"github.com/faridlan/firestore-vercel/repository"
	"github.com/faridlan/firestore-vercel/service"
	"github.com/gofiber/adaptor/v2"
)

func Handler(w http.ResponseWriter, r *http.Request) {

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

	adaptor.FiberApp(app)(w, r)

}
