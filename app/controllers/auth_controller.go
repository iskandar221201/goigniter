package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iskandar221201/goigniter/app/services"
	"github.com/iskandar221201/goigniter/app/validations"
	"github.com/iskandar221201/goigniter/system"
)

type AuthController struct {
	system.BaseController
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (ac *AuthController) Register(c *fiber.Ctx) error {
	var input validations.RegisterInput
	if err := ac.BodyParse(c, &input); err != nil {
		return err
	}

	user, err := ac.authService.Register(input)
	if err != nil {
		return ac.RespondError(c, fiber.StatusUnprocessableEntity, err.Error())
	}

	return ac.Respond(c, fiber.StatusCreated, "Registered", user)
}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	var input validations.LoginInput
	if err := ac.BodyParse(c, &input); err != nil {
		return err
	}

	token, err := ac.authService.Login(input)
	if err != nil {
		return ac.RespondError(c, fiber.StatusUnauthorized, err.Error())
	}

	return ac.Respond(c, fiber.StatusOK, "OK", fiber.Map{"token": token})
}
