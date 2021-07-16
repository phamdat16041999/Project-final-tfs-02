package model

import (
	"encoding/json"
	"fmt"
	"hotel/connect"
	"net/http"
	"strconv"
	"time"

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
type CreateBill struct {
	UserID    uint64    `json:"userID,omitempty"`
	HotelID   uint      `json:"hotelID,omitempty"`
	RoomID    uint      `json:"roomID,omitempty"`
	StartTime time.Time `json:"startTime,omitempty"`
	EndTime   time.Time `json:"endTime,omitempty"`
	Total     int       `json:"total,omitempty"`
}

func Createbill(w http.ResponseWriter, r *http.Request) {
	db := connect.Connect()
	userId := r.Context().Value("user_id")
	var createBill CreateBill
	var time Times
	var bill Bill
	err := json.NewDecoder(r.Body).Decode(&createBill)
	if err != nil {
		fmt.Fprint(w, err)
	}
	// convert interface to string
	str := fmt.Sprintf("%v", userId)
	// convert string str to unit to update in struct
	userid, _ := strconv.ParseUint(str, 10, 64)
	// b, err := json.Marshal(bill)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	fmt.Fprintln(w, userid)
	// fmt.Fprintln(w, string(b))
	// Tao time
	time.StartTime = createBill.StartTime
	time.EndTime = createBill.EndTime
	time.RoomID = createBill.RoomID
	result := db.Create(&time)
	if result.Error != nil {
		fmt.Fprintf(w, "Create time have error: %v", result)
		return
	} else {
		db.Last(&time)
		bill.UserID = userid
		bill.HotelID = createBill.HotelID
		bill.RoomID = createBill.RoomID
		bill.TimeID = time.ID
		bill.Total = createBill.Total
		result := db.Create(&bill)
		if result.Error != nil {
			fmt.Fprintf(w, "Create bill have error: %v", result)
			return
		} else {
			fmt.Fprintf(w, "Create bill have successfully")
		}

	}

	// result := db.Create(&bill)
}
func GetBill(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "aaaaaaaaaa")
}
