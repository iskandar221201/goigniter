// Package validations mendefinisikan struct input request dengan tag validate dari go-playground/validator.
package validations

type CreateUserInput struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Role     string `json:"role" validate:"omitempty,oneof=admin user"`
}

type UpdateUserInput struct {
	Name string `json:"name" validate:"omitempty,min=2,max=100"`
	Role string `json:"role" validate:"omitempty,oneof=admin user"`
}
