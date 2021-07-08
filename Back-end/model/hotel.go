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
	Name        string       `json:"authentication omitempty" gorm:"type:varchar(100);" json:"name" `
	Address     string       ` gorm:"type:varchar(100);" json:"address" `
	Description string       `gorm:"type:varchar(100);" json:"description" `
	Image       string       ` gorm:"type:varchar(100);" json:"image" `
	Longitude   string       ` gorm:"type:varchar(100);" json:"longitude" `
	Latitude    string       ` gorm:"type:varchar(100);" json:"latitude" `
	UserID      uint         `json:"userID"`
	ImageHotel  []ImageHotel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
	Room        []Room       `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
	Rate        []Rate       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
	Bill        []Bill       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
}

type ratehotel struct {
	HotelId uint
	Rate    int
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

func TopHotel(w http.ResponseWriter, r *http.Request) {
	db := connect.Connect()
	var rates []Rate
	db.Limit(2).Select("hotel_id", "rate").Order("rate desc").Find(&rates)
	w.Header().Set("Content-Type", "application/json")
	var rate1 []ratehotel
	for i := 0; i < len(rates); i++ {
		rate1 = append(rate1, ratehotel{HotelId: rates[i].HotelID,
			Rate: rates[i].Rate})
	}
	for i := 0; i < len(rate1); i++ {
		var hotels []Hotel
		db.Where("id = ?", rate1[i].HotelId).Find(&hotels)
		b, _ := json.Marshal(hotels)
		fmt.Fprintln(w, string(b))

	}

	b1, _ := json.Marshal(rate1)
	fmt.Fprintln(w, rate1[0].HotelId)
	fmt.Fprintln(w, string(b1))
}

func GetEachHotel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := connect.Connect()
	vars := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(vars)
	var query Hotel
	db.Where("id = ?", id).Find(&query)
	b, _ := json.Marshal(query)
	fmt.Fprintln(w, string(b))
}
