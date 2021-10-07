package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	Username string
	Password string
	IsAvailable bool
	Token    string  `gorm:"-" json:"token,omitempty"`
}

type CreateUser struct {
	Name string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	IsAvailable bool `json:"isAvailable"`
}

type UpdateUser struct {
	Name string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	IsAvailable bool `json:"isAvailable"`
}
