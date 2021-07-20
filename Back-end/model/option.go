package model

import (
	"encoding/json"
	"fmt"
	"hotel/connect"
	"net/http"

	"gorm.io/gorm"
)

type Option struct {
	gorm.Model
	Name  string  `gorm:"type:varchar(20);" json:"name,omitempty"`
	Price []Price `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:OptionID;associationForeignKey:ID"`
}

func OptionforHotel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := connect.Connect()
	var option []Option
	db.Debug().Find(&option)
	b, _ := json.Marshal(&option)
	fmt.Fprint(w, string(b))

}
