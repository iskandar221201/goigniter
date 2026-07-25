package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iskandar221201/goigniter/app/services"
	"github.com/iskandar221201/goigniter/app/validations"
	"github.com/iskandar221201/goigniter/system"
)

type UserController struct {
	system.BaseController
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) Index(c *fiber.Ctx) error {
	users, err := uc.userService.GetAll()
	if err != nil {
		return uc.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}

	return uc.Respond(c, fiber.StatusOK, "OK", users)
}

func (uc *UserController) Show(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return uc.RespondError(c, fiber.StatusBadRequest, "Invalid ID")
	}

	user, err := uc.userService.GetByID(uint(id))
	if err != nil {
		return uc.RespondError(c, fiber.StatusNotFound, "User not found")
	}

	return uc.Respond(c, fiber.StatusOK, "OK", user)
}

func (uc *UserController) Create(c *fiber.Ctx) error {
	var input validations.CreateUserInput
	if err := uc.BodyParse(c, &input); err != nil {
		return err
	}

	user, err := uc.userService.Create(input)
	if err != nil {
		return uc.RespondError(c, fiber.StatusUnprocessableEntity, err.Error())
	}

	return uc.Respond(c, fiber.StatusCreated, "Created", user)
}

func (uc *UserController) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return uc.RespondError(c, fiber.StatusBadRequest, "Invalid ID")
	}

	var input validations.UpdateUserInput
	if err := uc.BodyParse(c, &input); err != nil {
		return err
	}

	user, err := uc.userService.Update(uint(id), input)
	if err != nil {
		return uc.RespondError(c, fiber.StatusNotFound, err.Error())
	}

	return uc.Respond(c, fiber.StatusOK, "Updated", user)
}

func (uc *UserController) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return uc.RespondError(c, fiber.StatusBadRequest, "Invalid ID")
	}

	if err := uc.userService.Delete(uint(id)); err != nil {
		return uc.RespondError(c, fiber.StatusNotFound, err.Error())
	}

	return uc.Respond(c, fiber.StatusOK, "Deleted", nil)
}
