package model

import "gorm.io/gorm"

type Option struct {
	gorm.Model
	Name  string  `gorm:"type:varchar(20);" json:"name,omitempty"`
	Price []Price `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:OptionID;associationForeignKey:ID"`
}
