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

// func Payment(w http.ResponseWriter, r *http.Request) {
// 	var clientID = "AVbxI4aUPDOalr30IhZ5v-G58BIa1kUnL9pNjesbUFDlUf20GIdXQN7fXh2IydLrHRzjylo13LCVi1Vi"
// 	var secretID = "EPYJjuFqFyXD_84rhY-IQzi4s5FqTIrsdjVlT0m0rgo_iV_q-uknpLXKgWs3WVeqssMUjVN9FFAhkFh2"
// 	c, err := paypal.NewClient(clientID, secretID, paypal.APIBaseSandBox)
// 	if err != nil {
// 		fmt.Println("Error: ", err)
// 	}
// 	_, err = c.GetAccessToken()
// 	if err != nil {
// 		fmt.Println("Error: ", err)
// 	}

// 	var amount = "20"
// 	purchaseUnits := make([]paypal.PurchaseUnitRequest, 1)
// 	purchaseUnits[0] = paypal.PurchaseUnitRequest{
// 		Amount: &paypal.PurchaseUnitAmount{
// 			Currency: "USD",  //Payment type
// 			Value:    amount, //Received amount
// 		},
// 	}
// 	payer := &paypal.CreateOrderPayer{}
// 	appContext := &paypal.ApplicationContext{
// 		ReturnURL: "", //Callback link
// 	}
// 	order, err := c.CreateOrder("CAPTURE", purchaseUnits, payer, appContext)
// 	if err != nil {
// 		log.Error("create order errors:", err)
// 		fmt.Println("Error: ", err)
// 	}
// 	fmt.Println("Order: ", order)
// }
