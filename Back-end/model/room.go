package model

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	HotelID     uint        `json:"hotelID,omitempty"`
	ImageRoom   []ImageRoom `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:RoomID;associationForeignKey:ID"`
	Time        []Times     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:RoomID;associationForeignKey:ID"`
	Price       []Price     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:RoomID;associationForeignKey:ID"`
	Bill        []Bill      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:RoomID;associationForeignKey:ID"`
}
