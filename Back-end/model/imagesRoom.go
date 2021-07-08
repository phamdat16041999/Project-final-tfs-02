package model

import "gorm.io/gorm"

type ImageRoom struct {
	gorm.Model
	Image  string `json:"authentication omitempty" gorm:"type:varchar(100);" json:"image"`
	RoomID uint   `json:"roomID"`
}
