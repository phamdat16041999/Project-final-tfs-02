package model

import "gorm.io/gorm"

type ImageRoom struct {
	gorm.Model
	Image  string `gorm:"type:text;" json:"image,omitempty"`
	RoomID uint   `json:"roomID,omitempty"`
}
