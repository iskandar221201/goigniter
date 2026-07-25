package system

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type BaseController struct{}

func (bc *BaseController) Respond(c *fiber.Ctx, status int, message string, data interface{}) error {
	return Success(c, status, message, data)
}

func (bc *BaseController) RespondError(c *fiber.Ctx, status int, message string) error {
	return Error(c, status, message)
}

func (bc *BaseController) RespondValidation(c *fiber.Ctx, errors interface{}) error {
	return ValidationError(c, errors)
}

func (bc *BaseController) BodyParse(c *fiber.Ctx, out interface{}) error {
	if err := c.BodyParser(out); err != nil {
		return Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := validator.New().Struct(out); err != nil {
		errs := ParseValidationErrors(err)
		return ValidationError(c, errs)
	}

	return nil
}

func (bc *BaseController) CurrentUser(c *fiber.Ctx) *JWTClaims {
	user, ok := c.Locals("user").(*JWTClaims)
	if !ok {
		return nil
	}
	return user
}
