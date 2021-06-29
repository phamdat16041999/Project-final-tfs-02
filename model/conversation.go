package model

import "gorm.io/gorm"

type Conversation struct {
	gorm.Model
	User1ID   uint        `json:"user1ID"`
	User2ID   uint        `json:"user2ID"`
	Messenger []Messenger `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:ConversationID;associationForeignKey:ID"`
}
