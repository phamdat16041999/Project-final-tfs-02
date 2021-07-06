package model

import "gorm.io/gorm"

type Price struct {
	gorm.Model
	Price    int  `json:"price"`
	OptionID uint `json:"optionID"`
	RoomID   uint `json:"roomID"`
}
