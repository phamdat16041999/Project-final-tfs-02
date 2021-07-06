package model

import (
	"time"

	"gorm.io/gorm"
)

type Times struct {
	gorm.Model
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Date      time.Time `json:"date"`
	Active    *bool     `json:"active"`
	RoomID    uint      `json:"roomID"`
	Bill      []Bill    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:TimeID;associationForeignKey:ID"`
}
