package model

import (
	"encoding/json"
	"fmt"
	"hotel/connect"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Hotel struct {
	gorm.Model
	Name        string       `gorm:"type:varchar(100);" json:"name" `
	Address     string       `gorm:"type:varchar(100);" json:"address" `
	Description string       `gorm:"type:varchar(100);" json:"description" `
	Image       string       `gorm:"type:varchar(100);" json:"image" `
	Longitude   string       `gorm:"type:varchar(100);" json:"longitude" `
	Latitude    string       `gorm:"type:varchar(100);" json:"latitude" `
	UserID      uint         `json:"userID"`
	ImageHotel  []ImageHotel `json:"authentication omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
	Room        []Room       `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
	Rate        []Rate       `json:"authentication omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
	Bill        []Bill       `json:"authentication omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
}

type TopHotel struct {
	ID          int    `json:"id"`
	Name        string `gorm:"type:varchar(100);" json:"name" `
	Address     string `gorm:"type:varchar(100);" json:"address" `
	Description string `gorm:"type:varchar(100);" json:"description" `
	Image       string `gorm:"type:varchar(100);" json:"image" `
	Longitude   string `gorm:"type:varchar(100);" json:"longitude" `
	Latitude    string `gorm:"type:varchar(100);" json:"latitude" `
	UserID      uint   `json:"userID"`
	Rate        int    `json:"rate"`
}

// type Results struct {
// 	Address string `json:"address"`
// }

func DataHomePage(w http.ResponseWriter, r *http.Request) {
	db := connect.Connect()
	var results []Hotel
	db.Model(&Hotel{}).Distinct().Pluck("address", &results)
	b, _ := json.Marshal(results)
	fmt.Fprintln(w, string(b))

}
func GetHotelAddress(w http.ResponseWriter, r *http.Request) {
	db := connect.Connect()
	vars := mux.Vars(r)
	// vars["address"]
	var hotels []Hotel
	db.Where("address LIKE ?", vars["address"]).Find(&hotels)
	b, _ := json.Marshal(hotels)
	fmt.Fprintln(w, string(b))
}

func GetTopHotel(w http.ResponseWriter, r *http.Request) {
	db := connect.Connect()
	var rates []Rate
	db.Limit(2).Select("hotel_id", "rate").Order("rate desc").Find(&rates)
	w.Header().Set("Content-Type", "application/json")
	var topHotel []TopHotel
	var hotel Hotel
	for i := 0; i < len(rates); i++ {
		// b, _ := json.Marshal(rates[i].HotelID)
		// ID, _ := strconv.Atoi(string(b))
		// fmt.Fprintln(w, ID)
		db.First(&hotel, rates[i].HotelID)
		// b2, _ := json.Marshal(hotel)
		// fmt.Fprintln(w, hotel)
		v := TopHotel{
			ID:          int(hotel.ID),
			Name:        hotel.Name,
			Address:     hotel.Address,
			Description: hotel.Description,
			Image:       hotel.Image,
			Longitude:   hotel.Longitude,
			Latitude:    hotel.Latitude,
			Rate:        rates[i].Rate,
		}
		topHotel = append(topHotel, v)
		// topHotel = append(topHotel, hotel)

	}
	b2, _ := json.Marshal(topHotel)
	fmt.Fprintln(w, string(b2))

	// b1, _ := json.Marshal(rates)
	// fmt.Fprintln(w, rate1[0].HotelId)
	// fmt.Fprintln(w, string(b1))
}

func GetDetailHotel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := connect.Connect()
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var hotel Hotel
	var rooms []Room
	db.Where("id = ?", id).Find(&hotel)
	db.Where("hotel_id = ?", id).Find(&rooms)
	for i := 0; i < len(rooms); i++ {
		hotel.Room = append(hotel.Room, rooms[i])
		var imageRooms []ImageRoom
		db.Where("room_id = ?", rooms[i].ID).Find(&imageRooms)
		for j := 0; j < len(imageRooms); j++ {
			hotel.Room[i].ImageRoom = append(hotel.Room[i].ImageRoom, imageRooms[j])
		}
		var priceroom []Price
		db.Where("room_id = ?", rooms[i].ID).Find(&priceroom)
		for j := 0; j < len(priceroom); j++ {
			hotel.Room[i].Price = append(hotel.Room[i].Price, priceroom[j])
			// var option Option
			// db.Where("id = ?", priceroom[i].OptionID).Find(&option)
			// hotel.Room[i].Price[j].OptionID = append(hotel.Room[i].Price[j].OptionID, option)
		}
	}
	b1, _ := json.Marshal(hotel)
	fmt.Fprintln(w, string(b1))

}
