// Package models mendefinisikan struktur tabel database via GORM beserta method query-nya.
package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(255);not null"`
	Email    string `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Password string `json:"-"`
	Role     string `json:"role" gorm:"type:varchar(50);default:user"`
}

func (u *User) FindByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) FindByID(db *gorm.DB, id uint) (*User, error) {
	var user User
	err := db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) FindAll(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
