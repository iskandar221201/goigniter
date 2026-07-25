// Package services berisi business logic aplikasi, dipisahkan dari controller dan model.
package services

import (
	"github.com/iskandar221201/goigniter/app/helpers"
	"github.com/iskandar221201/goigniter/app/models"
	"github.com/iskandar221201/goigniter/app/validations"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) GetAll() ([]models.User, error) {
	users, err := new(models.User).FindAll(s.db)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetByID(id uint) (*models.User, error) {
	user, err := new(models.User).FindByID(s.db, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Create(input validations.CreateUserInput) (*models.User, error) {
	hash, err := helpers.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hash,
	}

	if input.Role != "" {
		user.Role = input.Role
	}

	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Update(id uint, input validations.UpdateUserInput) (*models.User, error) {
	user, err := new(models.User).FindByID(s.db, id)
	if err != nil {
		return nil, err
	}

	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Role != "" {
		user.Role = input.Role
	}

	if err := s.db.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Delete(id uint) error {
	user, err := new(models.User).FindByID(s.db, id)
	if err != nil {
		return err
	}

	return s.db.Delete(user).Error
}
