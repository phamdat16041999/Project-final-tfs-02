package model

import "gorm.io/gorm"

type Messenger struct {
	gorm.Model
	UserID         uint   `json:"userID"`
	Messenger      string `json:"authentication omitempty" gorm:"type:text;" json:"messenger"`
	ConversationID uint   `json:"conversationID"`
}
