package services

import (
	"errors"

	"github.com/iskandar221201/goigniter/app/helpers"
	"github.com/iskandar221201/goigniter/app/models"
	"github.com/iskandar221201/goigniter/app/validations"
	"github.com/iskandar221201/goigniter/config"
	"github.com/iskandar221201/goigniter/system"
	"gorm.io/gorm"
)

type AuthService struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewAuthService(db *gorm.DB, cfg *config.Config) *AuthService {
	return &AuthService{db: db, cfg: cfg}
}

func (s *AuthService) Register(input validations.RegisterInput) (*models.User, error) {
	hash, err := helpers.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hash,
	}

	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(input validations.LoginInput) (string, error) {
	var user models.User
	if err := s.db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid credentials")
		}
		return "", err
	}

	if !helpers.CheckPassword(input.Password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	claims := system.JWTClaims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
	}

	token, err := system.GenerateToken(claims, s.cfg)
	if err != nil {
		return "", err
	}

	return token, nil
}
