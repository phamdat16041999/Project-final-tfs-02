package model

import (
	"time"

	"gorm.io/gorm"
)

type Times struct {
	gorm.Model
	StartTime time.Time `json:"startTime,omitempty"`
	EndTime   time.Time `json:"endTime,omitempty"`
	Date      time.Time `json:"date,omitempty"`
	Active    *bool     `json:"active,omitempty"`
	RoomID    uint      `json:"roomID,omitempty"`
	Bill      []Bill    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:TimeID;associationForeignKey:ID"`
}
