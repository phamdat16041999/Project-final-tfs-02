package model

import "gorm.io/gorm"

type Option struct {
	gorm.Model
	Name  string  `gorm:"type:varchar(20);" json:"name"`
	Price []Price `json:"authentication omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:OptionID;associationForeignKey:ID"`
}
