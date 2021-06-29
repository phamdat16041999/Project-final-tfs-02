package model

import "gorm.io/gorm"

type Bill struct {
	gorm.Model
	UserID  uint `json:"userID"`
	HotelID uint `json:"hotelID"`
	RoomID  uint `json:"roomID"`
	TimeID  uint `json:"timeID"`
	Total   int  `json:"totalID"`
}
