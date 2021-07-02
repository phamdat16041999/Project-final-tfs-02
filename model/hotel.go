package model

import "gorm.io/gorm"

type Hotel struct {
	gorm.Model
	Name        string       `gorm:"type:varchar(100);" json:"name"`
	Address     string       `gorm:"type:varchar(100);" json:"address"`
	Description string       `gorm:"type:varchar(100);" json:"description"`
	Image       string       `gorm:"type:varchar(100);" json:"image"`
	Longitude   string       `gorm:"type:varchar(100);" json:"longitude"`
	Latitude    string       `gorm:"type:varchar(100);" json:"latitude"`
	UserID      uint         `json:"userID"`
	ImageHotel  []ImageHotel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
	Room        []Room       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
	Rate        []Rate       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
	Bill        []Bill       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
}
