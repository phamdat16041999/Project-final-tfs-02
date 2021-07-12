package model

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type Bill struct {
	gorm.Model
	UserID  uint `json:"userID"`
	HotelID uint `json:"hotelID"`
	RoomID  uint `json:"roomID"`
	TimeID  uint `json:"timeID"`
	Total   int  `json:"totalID"`
}
type DataToken struct {
	Authorized bool `json:"authorized"`
	Exp        int  `json:"exp"`
	Roles_id   int  `json:"roles_id"`
	User_id    int  `json:"user_id"`
}

func Createbill(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("user_id")
	// string --> fill struct
	// tk := make(jwt.MapClaims)

	// a, _ := json.Marshal(token)
	// var b DataToken
	// err := json.Unmarshal(a, &b)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	fmt.Fprint(w, token)

}
