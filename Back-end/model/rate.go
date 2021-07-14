package model

import "gorm.io/gorm"

type Rate struct {
	gorm.Model
	UserID  uint `json:"userID,omitempty"`
	HotelID uint `json:"hotelID,omitempty"`
	Rate    int  `json:"rate,omitempty"`
}
