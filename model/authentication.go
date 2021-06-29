package model

import (
	"gorm.io/gorm"
)

type Authentication struct {
	gorm.Model
	UserID uint `json:"userID"`
	RoleID uint `json:"roleID"`
}
