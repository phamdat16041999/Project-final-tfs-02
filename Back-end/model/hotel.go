package model

import (
	"encoding/json"
	"fmt"
	"hotel/connect"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Hotel struct {
	gorm.Model
	Name        string       `gorm:"type:varchar(100);" json:"name"`
	Address     string       `gorm:"type:varchar(100);" json:"address"`
	Description string       `gorm:"type:varchar(100);" json:"description"`
	Image       string       `gorm:"type:varchar(100);" json:"image"`
	Longitude   string       `gorm:"type:varchar(100);" json:"longitude"`
	Latitude    string       `gorm:"type:varchar(100);" json:"latitude"`
	UserID      uint         `json:"userID"`
	ImageHotel  []ImageHotel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
	Room        []Room       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
	Rate        []Rate       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
	Bill        []Bill       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:HotelID;associationForeignKey:ID"`
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
func GetHotel(w http.ResponseWriter, r *http.Request) {
	db := connect.Connect()
	vars := mux.Vars(r)
	// vars["address"]
	var hotels []Hotel
	db.Where("address LIKE ?", vars["address"]).Find(&hotels)
	b, _ := json.Marshal(hotels)
	fmt.Fprintln(w, string(b))
}
func TopHotel(w http.ResponseWriter, r *http.Request) {
	// db := connect.Connect()

}
