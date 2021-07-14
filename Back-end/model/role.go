package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name           string           `gorm:"type:varchar(100);" json:"name,omitempty"`
	Description    string           `gorm:"type:varchar(100);" json:"description,omitempty"`
	Authentication []Authentication `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:RoleID;associationForeignKey:ID"`
}
