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
	Room        []Room       `json:"room" gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
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
	db.Debug().Joins("JOIN rates on rates.hotel_id = hotels.id AND rates.rate = ?", vars["rate"]).Where("address LIKE ?", vars["address"]).Find(&hotels)
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
		b, _ := json.Marshal(rates[i].HotelID)
		ID, _ := strconv.Atoi(string(b))
		db.Where("id = ? ", ID).Find(&hotel)
		v := TopHotel{
			ID:          int(ID),
			Name:        hotel.Name,
			Address:     hotel.Address,
			Description: hotel.Description,
			Image:       hotel.Image,
			Longitude:   hotel.Longitude,
			Latitude:    hotel.Latitude,
			Rate:        rates[i].Rate,
		}
		topHotel = append(topHotel, v)
	}
	b2, _ := json.Marshal(topHotel)
	fmt.Fprintln(w, string(b2))
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
func CreateHotel(w http.ResponseWriter, r *http.Request) {
	db := connect.Connect()
	w.Header().Set("Content-Type", "application/json")
	userid := r.Context().Value("user_id")
	str := fmt.Sprintf("%v", userid)
	userID, _ := strconv.ParseUint(str, 10, 64)

	var Data Hotel
	err := json.NewDecoder(r.Body).Decode(&Data)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	hotel := Hotel{
		Name:        Data.Name,
		Address:     Data.Address,
		Description: Data.Description,
		Image:       Data.Image,
		Longitude:   Data.Longitude,
		Latitude:    Data.Latitude,
		UserID:      uint(userID),
	}
	result := db.Create(&hotel)
	if result.Error != nil {
		fmt.Fprint(w, result.Error)
	}
	// create each room
	for i := 0; i < len(Data.Room); i++ {
		room := Room{
			Name:        Data.Room[i].Name,
			Description: Data.Room[i].Description,
			HotelID:     hotel.ID,
		}
		result := db.Create(&room)
		if result.Error != nil {
			fmt.Fprint(w, result.Error)
			continue
		}
		// create each image for each room
		for j := 0; j < len(Data.Room[i].ImageRoom); j++ {
			imagerRoom := ImageRoom{
				Image:  Data.Room[i].ImageRoom[j].Image,
				RoomID: room.ID,
			}
			result := db.Create(&imagerRoom)
			if result.Error != nil {
				fmt.Fprint(w, result.Error)
				continue
			}
		}
		for k := 0; k < len(Data.Room[i].Price); k++ {
			price := Price{
				Price:    Data.Room[i].Price[k].Price,
				OptionID: Data.Room[i].Price[k].OptionID,
				RoomID:   room.ID,
			}
			result := db.Create(&price)
			if result.Error != nil {
				fmt.Fprint(w, result.Error)
				continue
			}

		}

	}
}
