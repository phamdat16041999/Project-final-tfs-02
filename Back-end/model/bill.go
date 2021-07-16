package model

import (
	"encoding/json"
	"fmt"
	"hotel/connect"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type Bill struct {
	gorm.Model
	UserID  uint64 `json:"userID,omitempty"`
	HotelID uint   `json:"hotelID,omitempty"`
	RoomID  uint   `json:"roomID,omitempty"`
	TimeID  uint   `json:"timeID,omitempty"`
	Total   int    `json:"total,omitempty"`
}

func Createbill(w http.ResponseWriter, r *http.Request) {
	db := connect.Connect()
	userId := r.Context().Value("user_id")
	var bill Bill
	err := json.NewDecoder(r.Body).Decode(&bill)
	if err != nil {
		fmt.Fprint(w, err)
	}
	// convert interface to string
	str := fmt.Sprintf("%v", userId)
	// convert string str to unit to update in struct
	userid, _ := strconv.ParseUint(str, 10, 64)
	bill.UserID = userid
	result := db.Create(&bill)
	if result.Error != nil {
		fmt.Fprintf(w, "Create bill have error: %v", bill)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Create bill successfull")
	}
}
func GetBill(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "aaaaaaaaaa")
}
