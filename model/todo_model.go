package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title string `json:"title" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}
