package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName          string           `gorm:"type:varchar(100); json:"firstName"`
	LastName           string           `gorm:"type:varchar(100); json:"lastName"`
	DOB                time.Time        `json:"dob"`
	Address            string           `gorm:"type:varchar(100); json:"address"`
	Phone              int              `json:"phone"`
	Email              string           `gorm:"type:varchar(100); json:"email"`
	CodeAuthentication string           `gorm:"type:varchar(20); json:"codeAuthentication"`
	UserName           string           `gorm:"type:varchar(100); json:"userName"`
	Password           string           `gorm:"type:varchar(100); json:"password"`
	Active             *bool            `json:"active"`
	Authentication     []Authentication `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:UserID;associationForeignKey:ID"`
	Conversation1      []Conversation   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:User1ID; associationForeignKey:ID"`
	Conversation2      []Conversation   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:User2ID; associationForeignKey:ID"`
	Messenger          []Messenger      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:UserID;associationForeignKey:ID"`
	Hotel              []Hotel          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:UserID;associationForeignKey:ID"`
	Rate               []Rate           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:UserID;associationForeignKey:ID"`
	Bill               []Bill           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:UserID;associationForeignKey:ID"`
}
