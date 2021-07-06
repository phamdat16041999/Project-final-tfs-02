package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name           string           `gorm:"type:varchar(100);" json:"name"`
	Description    string           `gorm:"type:varchar(100);" json:"description"`
	Authentication []Authentication `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:RoleID;associationForeignKey:ID"`
}
