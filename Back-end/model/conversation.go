package model

import (
	"hotel/connect"
	"strconv"

	"gorm.io/gorm"
)

type Conversation struct {
	gorm.Model
	User1ID   uint        `json:"user1ID"`
	User2ID   uint        `json:"user2ID"`
	Messenger []Messenger `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:ConversationID;associationForeignKey:ID"`
}

func CheckConvsersation(userID1, userID2 string) uint {
	db := connect.Connect()
	id1, _ := strconv.ParseUint(userID1, 10, 64)
	id2, _ := strconv.ParseUint(userID2, 10, 64)
	var checkConv1 Conversation
	var checkConv2 Conversation
	db.Where("user1_id = ? AND user2_id >= ?", id1, id2).Find(&checkConv1)
	db.Where("user1_id = ? AND user2_id >= ?", id2, id1).Find(&checkConv2)
	if checkConv1.ID == 0 && checkConv2.ID == 0 {
		covsversati1 := Conversation{
			User1ID: uint(id1),
			User2ID: uint(id2),
		}
		db.Create(&covsversati1)
		return covsversati1.ID
	} else {
		if checkConv1.ID == 0 {
			return checkConv2.ID
		} else {
			return checkConv1.ID
		}

	}

}
