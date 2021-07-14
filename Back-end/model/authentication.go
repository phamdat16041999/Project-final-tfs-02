package model

import (
	"gorm.io/gorm"
)

type Authentication struct {
	gorm.Model
	UserID uint `json:"userID,omitempty"`
	RoleID uint `json:"roleID,omitempty"`
}
