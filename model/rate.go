package model

import "gorm.io/gorm"

type Rate struct {
	gorm.Model
	UserID  uint `json:"userID"`
	HotelID uint `json:"hotelID"`
	Rate    int  `json:"rate"`
}
