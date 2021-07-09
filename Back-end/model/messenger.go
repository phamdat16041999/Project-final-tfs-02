package model

import "gorm.io/gorm"

type Messenger struct {
	gorm.Model
	UserID         uint   `json:"userID"`
	Messenger      string `gorm:"type:text;" json:"messenger"`
	ConversationID uint   `json:"conversationID"`
}
