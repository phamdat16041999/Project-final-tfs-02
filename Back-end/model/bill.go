package model

import (
	"encoding/json"
	"fmt"
	"hotel/connect"
	"hotel/middlewares"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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
type ShowBill struct {
	FirstName string    `json:"firstName,omitempty"`
	LastName  string    `json:"lastName,omitempty"`
	Address   string    `json:"address,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Email     string    `json:"email,omitempty"`
	Hotel     string    `json:"hotel,omitempty"`
	HotelID   uint      `json:"hotelID,omitempty"`
	Room      string    `json:"room,omitempty"`
	StartTime time.Time `json:"startTime,omitempty"`
	EndTime   time.Time `json:"endTime,omitempty"`
	Total     int       `json:"total,omitempty"`
}

func Createbill(w http.ResponseWriter, r *http.Request) {
	db := connect.Connect()
	data := r.Context().Value("data")
	UserID := middlewares.ConvertDataToken(data, "user_id")
	userid, err1 := strconv.ParseUint(UserID, 10, 64)
	if err1 != nil {
		fmt.Println("error:", err1)
	}
	var createBill CreateBill
	var time Times
	var bill Bill
	var showBill ShowBill
	var user User
	var hotel Hotel
	var room Room
	err := json.NewDecoder(r.Body).Decode(&createBill)
	if err != nil {
		fmt.Fprint(w, err)
	}
	// convert string str to unit to update in struct
	fmt.Println(userid)
	// Tao time
	time.StartTime = createBill.StartTime
	time.EndTime = createBill.EndTime
	time.RoomID = createBill.RoomID
	result := db.Create(&time)
	if result.Error != nil {
		fmt.Fprintf(w, "Create time have error: %v", result)
	} else {
		db.Last(&time)
		bill.UserID = userid
		bill.HotelID = createBill.HotelID
		bill.RoomID = createBill.RoomID
		bill.TimeID = time.ID
		bill.Total = createBill.Total
		result := db.Debug().Create(&bill)
		if result.Error != nil {
			fmt.Fprintf(w, "Create bill have error: %v", result)

		} else {
			db.First(&user, "id = ?", userid)
			db.First(&hotel, "id = ?", createBill.HotelID)
			db.First(&room, "id = ?", createBill.RoomID)
			showBill.FirstName = user.FirstName
			showBill.LastName = user.LastName
			showBill.Address = user.Address
			showBill.Email = user.Email
			showBill.Phone = user.Phone
			showBill.Hotel = hotel.Name
			showBill.HotelID = createBill.HotelID
			showBill.Room = room.Name
			showBill.StartTime = createBill.StartTime
			showBill.EndTime = createBill.EndTime
			showBill.Total = createBill.Total
			b, err := json.Marshal(showBill)
			if err != nil {
				fmt.Println("error:", err)
			}
			fmt.Fprintf(w, string(b))
		}

	}

	// result := db.Create(&bill)
}

type ListBill struct {
	NameHotel  string    `json:"namehotel"`
	NameRoom   string    `json:"nameroom"`
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
	TotalPrice int
}

func GetBill(w http.ResponseWriter, r *http.Request) {
	db := connect.Connect()
	data := r.Context().Value("data")
	UserID := middlewares.ConvertDataToken(data, "user_id")
	userid, err1 := strconv.ParseUint(UserID, 10, 64)
	if err1 != nil {
		fmt.Println("error:", err1)
	}
	var bills []Listbillofmanager
	var listbill []Bill
	var user User
	db.Where("id = ?", userid).Find(&user)
	db.Where("user_id =?", userid).Find(&listbill)
	for i := 0; i < len(listbill); i++ {
		var hotel Hotel
		var roomID Room
		var time Times
		db.Where("id = ?", listbill[i].HotelID).Find(&hotel)
		db.Where("id = ?", listbill[i].RoomID).Find(&roomID)
		db.Where("id = ?", listbill[i].TimeID).Find(&time)
		bills = append(bills, Listbillofmanager{
			ID:           int(listbill[i].ID),
			NameCustomer: user.FirstName + user.LastName,
			Address:      user.Address,
			Mail:         user.Email,
			Phone:        user.Phone,
			NameHotel:    hotel.Name,
			NameRoom:     roomID.Name,
			StartTime:    time.StartTime,
			EndTime:      time.EndTime,
			TotalPrice:   listbill[i].Total,
		})
	}
	b, err := json.Marshal(&bills)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Fprint(w, string(b))
}
func Detailbill(w http.ResponseWriter, r *http.Request) {
	db := connect.Connect()
	idbill, _ := strconv.Atoi(mux.Vars(r)["id"])
	var bill Bill
	db.Where("id =?", idbill).Find(&bill)
	var hotel Hotel
	var room Room
	var time Times
	var showBill ShowBill
	var user User
	db.Where("id = ?", bill.HotelID).Find(&hotel)
	db.Where("id = ?", bill.RoomID).Find(&room)
	db.Where("id = ?", bill.TimeID).Find(&time)
	db.Where("id = ?", bill.UserID).Find(&user)
	showBill.FirstName = user.FirstName
	showBill.LastName = user.LastName
	showBill.Address = user.Address
	showBill.Phone = user.Phone
	showBill.Email = user.Email
	showBill.Hotel = hotel.Name
	showBill.Room = room.Name
	showBill.StartTime = time.StartTime
	showBill.EndTime = time.EndTime
	showBill.Total = bill.Total

	b, err := json.Marshal(&showBill)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Fprint(w, string(b))
}

type Listbillofmanager struct {
	ID           int       `json:"id"`
	NameCustomer string    `json:"nameCustomer"`
	Address      string    `json:"address"`
	Mail         string    `json:"mail"`
	Phone        string    `json:"phone"`
	NameHotel    string    `json:"namehotel"`
	NameRoom     string    `json:"nameroom"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
	TotalPrice   int
}

func Allbillofmanagerhotel(w http.ResponseWriter, r *http.Request) {
	db := connect.Connect()
	data := r.Context().Value("data")
	UserID := middlewares.ConvertDataToken(data, "user_id")
	userid, err1 := strconv.ParseUint(UserID, 10, 64)
	if err1 != nil {
		fmt.Println("error:", err1)
	}
	var listbillofmanager []Listbillofmanager
	var hotels []Hotel
	var bill []Bill
	db.Debug().Where("user_id = ?", userid).Find(&hotels)
	fmt.Print(hotels)
	for _, hotel := range hotels {
		db.Where("hotel_id =?", hotel.ID).Find(&bill)
		for _, bills := range bill {
			var roomID Room
			var time Times
			var customer User
			db.Where("id = ?", bills.RoomID).Find(&roomID)
			db.Where("id = ?", bills.TimeID).Find(&time)
			db.Where("id = ?", bills.UserID).Find(&customer)
			listbillofmanager = append(listbillofmanager, Listbillofmanager{
				ID:           int(bills.ID),
				NameCustomer: customer.FirstName + ` ` + customer.LastName,
				Address:      customer.Address,
				Mail:         customer.Email,
				Phone:        customer.Phone,
				NameHotel:    hotel.Name,
				NameRoom:     roomID.Name,
				StartTime:    time.StartTime,
				EndTime:      time.EndTime,
				TotalPrice:   bills.Total,
			})
		}
	}
	b, err := json.Marshal(&listbillofmanager)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Fprint(w, string(b))

}
