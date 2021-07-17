package model

import (
	"encoding/json"
	"fmt"
	"hotel/connect"
	"hotel/pkg"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Hotel struct {
	gorm.Model
	Name        string       `gorm:"type:varchar(100);" json:"name,omitempty" `
	Address     string       `gorm:"type:varchar(100);" json:"address,omitempty" `
	Description string       `gorm:"type:varchar(100);" json:"description,omitempty" `
	Image       string       `gorm:"type:varchar(100);" json:"image,omitempty" `
	Longitude   string       `gorm:"type:varchar(100);" json:"longitude,omitempty" `
	Latitude    string       `gorm:"type:varchar(100);" json:"latitude,omitempty" `
	UserID      uint         `json:"userID,omitempty"`
	AverageRate float64      `gorm:"default:0.0;" json:"averagerate,omitempty"`
	NumberRate  float64      `gorm:"default:0;" json:"numberrate,omitempty"`
	ImageHotel  []ImageHotel `json:"imagehotel,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
	Room        []Room       `json:"room,omitempty" gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
	Rate        []Rate       `json:"rate,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
	Bill        []Bill       `json:"bill,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
}

type TopHotel struct {
	ID          uint64 `json:"id"`
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

type HotelRate struct {
	HotelID uint `json:"hotelID"`
	UserID  uint `json:"userID"`
	Rate    int  `json:"rate"`
}

func DataHomePage(w http.ResponseWriter, r *http.Request) {
	db := connect.Connect()
	var results []Hotel
	db.Model(&Hotel{}).Distinct().Pluck("address", &results)
	b, _ := json.Marshal(results)
	fmt.Fprintln(w, string(b))
}

func SearchByName(w http.ResponseWriter, r *http.Request) {

}
func GetHotelAddress(w http.ResponseWriter, r *http.Request) {

	db := connect.Connect()
	vars := mux.Vars(r)
	rate, _ := strconv.ParseFloat(vars["rate"], 64)
	var hotels []Hotel
	var resulthotels []Hotel
	db.Where("address LIKE ?", "%"+vars["address"]+"%").Find(&hotels)
	for i := 0; i < len(hotels); i++ {
		if int64(hotels[i].AverageRate) == int64(rate) {
			resulthotels = append(resulthotels, hotels[i])
		}
	}
	b, _ := json.Marshal(resulthotels)
	fmt.Fprintln(w, string(b))
}

func GetTopHotel(w http.ResponseWriter, r *http.Request) {

	db := connect.Connect()
	var hotel []Hotel
	db.Limit(2).Order("average_rate desc").Find(&hotel)
	b, _ := json.Marshal(hotel)
	pkg.ServeJQueryWithCache(w, "tophotel", string(b))
}
func SearchHotelAddress(w http.ResponseWriter, r *http.Request) {

	db := connect.Connect()
	vars := mux.Vars(r)
	rate, _ := strconv.ParseFloat(vars["rate"], 64)
	var hotels []Hotel
	var resulthotels []Hotel
	db.Debug().Where("address LIKE ?", "%"+vars["address"]+"%").Find(&hotels)
	for i := 0; i < len(hotels); i++ {
		if int64(hotels[i].AverageRate) == int64(rate) {
			resulthotels = append(resulthotels, hotels[i])
		}
	}
	b, _ := json.Marshal(resulthotels)
	fmt.Fprintln(w, string(b))
}
func GetDetailHotel(w http.ResponseWriter, r *http.Request) {
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

func Rating(w http.ResponseWriter, r *http.Request) {
	pkg.DeleteRemoteCache(w, "tophotel")
	pkg.DeleteLocalCache(w, "tophotel")

	var hotelrate HotelRate
	var rate Rate
	var hotel Hotel
	err := json.NewDecoder(r.Body).Decode(&hotelrate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db := connect.Connect()
	db.Where("user_id = ? AND hotel_id = ?", hotelrate.UserID, hotelrate.HotelID).Find(&rate)
	b1, _ := json.Marshal(&rate.HotelID)
	ID, _ := strconv.ParseUint(string(b1), 10, 32)
	b, _ := json.Marshal(&rate.Rate)

	var Rate = Rate{
		UserID:  hotelrate.UserID,
		HotelID: hotelrate.HotelID,
		Rate:    hotelrate.Rate,
	}
	if string(b) == "0" {
		result := db.Create(&Rate)
		if result.Error != nil {
			fmt.Fprintln(w, "Rating error: ", result.Error)
			return
		} else {
			db.Where("id = ?", hotelrate.HotelID).Find(&hotel)
			b, _ := json.Marshal(&hotel.NumberRate)
			NumberRate, _ := strconv.ParseFloat(string(b), 64)
			b1, _ := json.Marshal(&hotel.AverageRate)
			b3, _ := json.Marshal(&hotelrate.Rate)
			Rate, _ := strconv.ParseUint(string(b3), 10, 64)
			AverageRate, _ := strconv.ParseFloat(string(b1), 64)
			AverageRate = (AverageRate*NumberRate + float64(Rate)) / (NumberRate + 1)
			db.Model(Hotel{}).Where("id = ?", hotelrate.HotelID).Updates(Hotel{NumberRate: NumberRate + 1.0, AverageRate: AverageRate})
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "Rating successfull")
		}
	} else {
		b3, _ := json.Marshal(&rate.Rate)
		result := db.Model(&rate).Where("user_id = ? AND hotel_id = ?", hotelrate.UserID, hotelrate.HotelID).Update("rate", hotelrate.Rate)
		db.Where("id = ?", ID).Find(&hotel)
		b, _ := json.Marshal(&hotel.NumberRate)
		NumberRate, _ := strconv.ParseFloat(string(b), 64)
		b1, _ := json.Marshal(&hotel.AverageRate)
		b2, _ := json.Marshal(&hotelrate.Rate)
		NewRate, _ := strconv.ParseUint(string(b2), 10, 64)
		OldRate, _ := strconv.ParseUint(string(b3), 10, 64)
		fmt.Fprintln(w, "NewRate = ", NewRate, "\nOldRate = ", OldRate)
		AverageRate, _ := strconv.ParseFloat(string(b1), 64)
		AverageRate = (AverageRate*NumberRate - float64(OldRate) + float64(NewRate)) / NumberRate
		AverageRate = float64(int(AverageRate*10)) / 10 //chuyển thành số thập phân có 2 chữ số
		fmt.Fprintln(w, "AverageRate = ", AverageRate)
		db.Model(Hotel{}).Where("id = ?", hotelrate.HotelID).Updates(Hotel{AverageRate: AverageRate})
		if result.Error != nil {
			fmt.Fprintln(w, "Update rating error: ", result.Error)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "Update rating successfull")
		}
	}
}
func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

// func Converts(time string) int {
// 	a := strings.Split(time, "T")
// 	x1 := strings.Split(a[1], ":")
// 	x2 := strings.Split(a[0], "-")
// 	var s1, s2 []int
// 	for i := 0; i < len(x1); i++ {
// 		x1, _ := strconv.Atoi(x1[i])
// 		s1 = append(s1, x1)
// 	}
// 	for j := 0; j < len(x2); j++ {
// 		x2, _ := strconv.Atoi(x2[j])
// 		s2 = append(s2, x2)
// 	}
// 	hashtime := (s1[0]-00)*3600 + (s1[1]-00)*60
// 	hashdate := (s2[1]-7)*30 + (s2[1]-01)*1
// 	hash := hashtime + hashdate
// 	return hash
// }

func Checkroomstatus(w http.ResponseWriter, r *http.Request) {

	var checkTime Times
	err := json.NewDecoder(r.Body).Decode(&checkTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db := connect.Connect()
	var times []Times
	result := db.Where("room_id = ?", checkTime.RoomID).Find(&times)
	if result.Error != nil {
		fmt.Fprintln(w, "Can not find room: ", result.Error)
		return
	}
	available := 1
	for i := 0; i < len(times); i++ {
		if inTimeSpan(times[i].StartTime, times[i].EndTime, checkTime.StartTime) {
			available--
			break
		} else if inTimeSpan(times[i].StartTime, times[i].EndTime, checkTime.EndTime) {
			available--
			break
		}
	}
	if available != 1 {
		fmt.Fprintln(w, "Room has been booked")
	} else {
		fmt.Fprintln(w, "Room avaliable")
	}
}
