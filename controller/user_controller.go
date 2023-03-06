package controller

import (
	"fmt"
	"net/http"

	"github.com/go-jwt-test/data/request"
	"github.com/go-jwt-test/data/response"
	"github.com/go-jwt-test/helper"
	"github.com/go-jwt-test/service"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{UserService: service}
}

func (controller *UserController) Create(ctx *fiber.Ctx) error {
	createUserRequest := request.CreateUserRequest{}
	err := ctx.BodyParser(&createUserRequest)
	helper.ErrorPanic(err)

	controller.UserService.Create(createUserRequest)

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok!",
		Message: "Successfully created user data!",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)

}

func (controller *UserController) Update(xtc *fiber.Ctx) error {
	update := request.UpdateUserRequest{}
	err := xtc.BodyParser(&update)
	helper.ErrorPanic(err)

	userId := xtc.Params("uid")
	id := fmt.Sprint(userId)

	update.Uid = id

	controller.UserService.Update(update)

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok!",
		Message: "successfully updated user data",
		Data:    nil,
	}
	return xtc.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *UserController) Delete(ctx *fiber.Ctx) error {
	userId := ctx.Params("uid")
	uid := fmt.Sprint(userId)

	controller.UserService.Delete(uid)

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok!",
		Message: "successfully deleted user data",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)

}

func (controller *UserController) FindById(ctx *fiber.Ctx) error {
	userId := ctx.Params("uid")
	id := fmt.Sprint(userId)

	userResp := controller.UserService.FindByid(id)

	userResponse := response.Response{
		Code:    200,
		Status:  "Ok!",
		Message: "successfully get user data by id",
		Data:    userResp,
	}
	return ctx.Status(fiber.StatusCreated).JSON(userResponse)
}

func (controller *UserController) FindAll(ctx *fiber.Ctx) error {
	userResp := controller.UserService.FindAll()
	userResponse := response.Response{
		Code:    200,
		Status:  "Ok!",
		Message: "successfully fetch all user",
		Data:    userResp,
	}
	return ctx.Status(fiber.StatusCreated).JSON(userResponse)
}
