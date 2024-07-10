package controller

import (
	"github.com/faridlan/firestore-vercel/model/web"
	"github.com/faridlan/firestore-vercel/service"
	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	Save(ctx *fiber.Ctx) error
	Find(ctx *fiber.Ctx) error
}

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Save(ctx *fiber.Ctx) error {

	user := new(web.UserWeb)
	err := ctx.BodyParser(user)
	if err != nil {
		return err
	}

	userResponse, err := controller.UserService.Save(ctx.Context(), user)
	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	return ctx.JSON(webResponse)

}

func (controller *UserControllerImpl) Find(ctx *fiber.Ctx) error {

	userResponse, err := controller.UserService.Find(ctx.Context())
	if err != nil {
		return err
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	return ctx.JSON(webResponse)

}
