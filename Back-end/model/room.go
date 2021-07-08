package model

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Name        string      `json:"name"`
	Description string      `json:"description"`
	HotelID     uint        `json:"hotelID"`
	ImageRoom   []ImageRoom ` json:"authentication omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:RoomID;associationForeignKey:ID"`
	Time        []Times     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:RoomID;associationForeignKey:ID"`
	Price       []Price     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:RoomID;associationForeignKey:ID"`
	Bill        []Bill      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:RoomID;associationForeignKey:ID"`
}
