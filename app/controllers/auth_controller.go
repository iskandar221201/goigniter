package controllers

import (
	"net/http"

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

func (ac *AuthController) Register(w http.ResponseWriter, r *http.Request) error {
	var input validations.RegisterInput
	if err := ac.BodyParse(w, r, &input); err != nil {
		return err
	}

	user, err := ac.authService.Register(input)
	if err != nil {
		return ac.RespondError(w, http.StatusUnprocessableEntity, err.Error())
	}

	return ac.Respond(w, http.StatusCreated, "Registered", user)
}

func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) error {
	var input validations.LoginInput
	if err := ac.BodyParse(w, r, &input); err != nil {
		return err
	}

	token, err := ac.authService.Login(input)
	if err != nil {
		return ac.RespondError(w, http.StatusUnauthorized, err.Error())
	}

	return ac.Respond(w, http.StatusOK, "OK", map[string]string{"token": token})
}
