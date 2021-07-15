package model

import "gorm.io/gorm"

type ImageHotel struct {
	gorm.Model
	Image   string `json:"image,omitempty"`
	HotelID uint   `json:"hotelID,omitempty"`
}
