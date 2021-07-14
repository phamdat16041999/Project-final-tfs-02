package model

import "gorm.io/gorm"

type ImageRoom struct {
	gorm.Model
	Image  string `gorm:"type:varchar(100);" json:"image,omitempty"`
	RoomID uint   `json:"roomID,omitempty"`
}
