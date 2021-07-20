package model

import (
	"hotel/connect"
	"strconv"

	"gorm.io/gorm"
)

type Messenger struct {
	gorm.Model
	UserID         uint   `json:"userID,omitempty"`
	Messenger      string `gorm:"type:text;" json:"messenger,omitempty"`
	ConversationID uint   `json:"conversationID,omitempty"`
}

func CreateMessenger(UserID, messenger string, conversationID uint) []Messenger {
	db := connect.Connect()
	id, _ := strconv.ParseUint(UserID, 10, 64)
	if messenger != "" {
		messeng := Messenger{
			UserID:         uint(id),
			Messenger:      messenger,
			ConversationID: conversationID,
		}
		db.Create(&messeng)
	}
	var mess []Messenger
	db.Where("conversation_id = ?", conversationID).Find(&mess)
	return mess
}
