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
type HotelInformation struct {
	ID          int               `json:"id"`
	Name        string            `json:"name"`
	Room        []RoomInformation `json:"room"`
	Longitude   string            `json:"longitude"`
	Latitude    string            `json:"latitude"`
	Description string            `json:"description"`
}
type RoomInformation struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Img        []ImageRoom
	PriceHrs   int `json:"priceHrs"`
	PriceDay   int `json:"priceDay"`
	ExtraPrice int `json:"extraPrice"`
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
	var imageRoom []ImageRoom

	db.Where("id = ?", id).Find(&hotel)
	// db.Where("hotel_id = ?", id).Find(&rooms)
	// for i := 0; i < len(rooms); i++ {
	// 	hotel.Room = append(hotel.Room, rooms[i])
	// 	var imageRooms []ImageRoom
	// 	db.Where("room_id = ?", rooms[i].ID).Find(&imageRooms)
	// 	for j := 0; j < len(imageRooms); j++ {
	// 		hotel.Room[i].ImageRoom = append(hotel.Room[i].ImageRoom, imageRooms[j])
	// 	}
	// 	var priceroom []Price
	// 	db.Where("room_id = ?", rooms[i].ID).Find(&priceroom)
	// 	for j := 0; j < len(priceroom); j++ {
	// 		hotel.Room[i].Price = append(hotel.Room[i].Price, priceroom[j])
	// 		// var option Option
	// 		// db.Where("id = ?", priceroom[i].OptionID).Find(&option)
	// 		// hotel.Room[i].Price[j].OptionID = append(hotel.Room[i].Price[j].OptionID, option)
	// 	}
	// }
	hotelInformation := HotelInformation{
		ID:   int(hotel.ID),
		Name: hotel.Name,
		// room       : []RoomInformation
		Longitude:   hotel.Latitude,
		Latitude:    hotel.Longitude,
		Description: hotel.Description,
	}
	db.Where("hotel_id = ?", hotel.ID).Find(&rooms)
	for i := 0; i < len(rooms); i++ {
		var option1 Option
		var option2 Option
		var option3 Option
		var priceHrs Price
		var priceDay Price
		var extraPrice Price
		db.Where("room_id = ?", rooms[i].ID).Find(&imageRoom)
		db.Where("name = ?", "PriceHours").Find(&option1)
		db.Where("room_id = ? AND option_id = ?", rooms[i].ID, option1.ID).Find(&priceHrs)
		db.Where("name = ?", "PriceDays").Find(&option2)
		db.Where("room_id = ? AND option_id = ?", rooms[i].ID, option2.ID).Find(&priceDay)
		db.Where("name = ?", "ExtraPrice").Find(&option3)
		db.Where("room_id = ? AND option_id = ?", rooms[i].ID, option3.ID).Find(&extraPrice)

		roomInformation := RoomInformation{
			ID:         int(rooms[i].ID),
			Name:       rooms[i].Name,
			Img:        imageRoom,
			PriceHrs:   priceHrs.Price,
			PriceDay:   priceDay.Price,
			ExtraPrice: extraPrice.Price,
		}
		hotelInformation.Room = append(hotelInformation.Room, roomInformation)
	}
	b1, _ := json.Marshal(hotelInformation)
	fmt.Fprintln(w, string(b1))

}
