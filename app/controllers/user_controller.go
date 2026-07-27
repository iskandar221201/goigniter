package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

func (uc *UserController) Index(w http.ResponseWriter, r *http.Request) error {
	users, err := uc.userService.GetAll()
	if err != nil {
		return uc.RespondError(w, http.StatusInternalServerError, err.Error())
	}

	return uc.Respond(w, http.StatusOK, "OK", users)
}

func (uc *UserController) Show(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return uc.RespondError(w, http.StatusBadRequest, "Invalid ID")
	}

	user, err := uc.userService.GetByID(uint(id))
	if err != nil {
		return uc.RespondError(w, http.StatusNotFound, "User not found")
	}

	return uc.Respond(w, http.StatusOK, "OK", user)
}

func (uc *UserController) Create(w http.ResponseWriter, r *http.Request) error {
	var input validations.CreateUserInput
	if err := uc.BodyParse(w, r, &input); err != nil {
		return err
	}

	user, err := uc.userService.Create(input)
	if err != nil {
		return uc.RespondError(w, http.StatusUnprocessableEntity, err.Error())
	}

	return uc.Respond(w, http.StatusCreated, "Created", user)
}

func (uc *UserController) Update(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return uc.RespondError(w, http.StatusBadRequest, "Invalid ID")
	}

	var input validations.UpdateUserInput
	if err := uc.BodyParse(w, r, &input); err != nil {
		return err
	}

	user, err := uc.userService.Update(uint(id), input)
	if err != nil {
		return uc.RespondError(w, http.StatusNotFound, err.Error())
	}

	return uc.Respond(w, http.StatusOK, "Updated", user)
}

func (uc *UserController) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return uc.RespondError(w, http.StatusBadRequest, "Invalid ID")
	}

	if err := uc.userService.Delete(uint(id)); err != nil {
		return uc.RespondError(w, http.StatusNotFound, err.Error())
	}

	return uc.Respond(w, http.StatusOK, "Deleted", nil)
}
