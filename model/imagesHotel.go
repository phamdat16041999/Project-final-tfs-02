package model

import "gorm.io/gorm"

type ImageHotel struct {
	gorm.Model
	Image   string `json:"image"`
	HotelID uint   `json:"hotelID"`
}
