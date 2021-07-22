package model

import "gorm.io/gorm"

type Price struct {
	gorm.Model
	Price    string `json:"price,omitempty"`
	OptionID uint   `json:"optionID,omitempty"`
	RoomID   uint   `json:"roomID,omitempty"`
}
