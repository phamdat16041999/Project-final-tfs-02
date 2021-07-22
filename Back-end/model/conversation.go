package model

import (
	"encoding/json"
	"fmt"
	"hotel/connect"
	"hotel/middlewares"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type Conversation struct {
	gorm.Model
	User1ID   uint        `json:"user1ID,omitempty"`
	User2ID   uint        `json:"user2ID,omitempty"`
	Messenger []Messenger `json:"messenger,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:ConversationID;associationForeignKey:ID"`
}

func CheckConvsersation(userID1, userID2 string) uint {
	db := connect.Connect()
	id1, _ := strconv.ParseUint(userID1, 10, 64)
	id2, _ := strconv.ParseUint(userID2, 10, 64)
	var checkConv1 Conversation
	var checkConv2 Conversation
	db.Where("user1_id = ? AND user2_id = ?", id1, id2).Find(&checkConv1)
	db.Where("user1_id = ? AND user2_id = ?", id2, id1).Find(&checkConv2)
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
func ListConversation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := connect.Connect()
	data := r.Context().Value("data")
	UserID := middlewares.ConvertDataToken(data, "user_id")
	userid, err1 := strconv.ParseUint(UserID, 10, 64)
	if err1 != nil {
		fmt.Println("error:", err1)
	}
	var conversation []Conversation
	db.Where("user1_id =? OR user2_id =?", userid, userid).Find(&conversation)
	var users []User
	for _, conversations := range conversation {
		if userid == uint64(conversations.User1ID) {
			var userss User
			db.Where("id =?", conversations.User2ID).Find(&userss)
			users = append(users, userss)
		} else {
			var userss User
			db.Where("id =?", conversations.User1ID).Find(&userss)
			users = append(users, userss)
		}
	}
	b, _ := json.Marshal(&users)
	fmt.Fprint(w, string(b))
}
