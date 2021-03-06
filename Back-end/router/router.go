package router

import (
	"encoding/json"
	"fmt"
	"hotel/connect"
	"hotel/middlewares"
	"hotel/model"
	"strconv"

	//"hotel/pkg"
	"hotel/pkg/websocket"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Run() {
	// go pkg.RunRmq()
	r := mux.NewRouter().StrictSlash(true)
	post := r.Methods(http.MethodPost).Subrouter()
	post.Path("/account").HandlerFunc(model.CreateAccount)
	post.Path("/login").HandlerFunc(model.LoginAcount)
	post.Path("/forgotpassword").HandlerFunc(model.ForgotPassword)
	post.Path("/changepassword").HandlerFunc(middlewares.SetMiddlewareAuthentication(model.ChangePassword))
	post.Path("/active").HandlerFunc(model.ActiveAccount)
	post.Path("/createbill").HandlerFunc(middlewares.SetMiddlewareAuthenticationUser(model.Createbill))
	post.Path("/rating").HandlerFunc(middlewares.SetMiddlewareAuthentication(model.Rating))
	post.Path("/checkroomstatus").HandlerFunc(middlewares.SetMiddlewareAuthenticationUser(model.Checkroomstatus))
	post.Path("/createhotel").HandlerFunc(middlewares.SetMiddlewareAuthenticationHotelOwner(model.CreateHotel))
	// post.Path("/payment").HandlerFunc(middlewares.SetMiddlewareAuthenticationUser(model.Payment))

	//get method
	get := r.Methods(http.MethodGet).Subrouter()
	get.Path("/homepage").HandlerFunc(model.DataHomePage)
	get.Path("/hotel/{address}/{rate}").HandlerFunc(model.SearchHotelAddress)
	get.Path("/tophotel").HandlerFunc(model.GetTopHotel)
	get.Path("/detailhotel/{id}").HandlerFunc(model.GetDetailHotel)
	get.Path("/getbill").HandlerFunc(middlewares.SetMiddlewareAuthenticationUser(model.GetBill))
	get.Path("/search/{name}").HandlerFunc(model.SearchByName)
	get.Path("/essearch/{name}").HandlerFunc(model.EsSearchByName)
	get.Path("/hotelier").HandlerFunc(middlewares.SetMiddlewareAuthenticationHotelOwner(model.Hotelier))
	get.Path("/option").HandlerFunc(model.OptionforHotel)
	get.Path("/detailbillofmanagerhotel").HandlerFunc(middlewares.SetMiddlewareAuthenticationHotelOwner(model.Allbillofmanagerhotel))
	get.Path("/detailbill/{id}").HandlerFunc(middlewares.SetMiddlewareAuthentication(model.Detailbill))
	get.Path("/listConversation").HandlerFunc(middlewares.SetMiddlewareAuthentication(model.ListConversation))

	// delete method
	detelte := r.Methods(http.MethodDelete).Subrouter()
	detelte.Path("/hotel/{id}").HandlerFunc(middlewares.SetMiddlewareAuthenticationHotelOwner(model.DeleteHotel))
	// methodput
	r.HandleFunc("/update/{id}", middlewares.SetMiddlewareAuthentication(model.UpdateAccount)).Methods("PUT")
	r.HandleFunc("/updatehotel", middlewares.SetMiddlewareAuthentication(model.UpdateHotel)).Methods("PUT")
	r.HandleFunc("/checklogin", middlewares.SetMiddlewareAuthentication(CheckLogin)).Methods("GET")
	http.Handle("/", r)
	//methoddelete
	r.HandleFunc("/delete/{id}", middlewares.SetMiddlewareAuthentication(model.DeleteAccount)).Methods("Delete")
	// chat API
	r.HandleFunc("/ws", websocket.HandleConnections)

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS", "PUT"},
		AllowedHeaders: []string{"*"},
	}).Handler(r)
	http.ListenAndServe(":8080", handler)
}

type CheckLogins struct {
	Role   string `json:"role"`
	Status string `json:"status"`
}

func CheckLogin(w http.ResponseWriter, r *http.Request) {
	db := connect.Connect()
	data := r.Context().Value("data")
	UserID := middlewares.ConvertDataToken(data, "user_id")
	userid, err1 := strconv.ParseUint(UserID, 10, 64)
	if err1 != nil {
		fmt.Println("error:", err1)
	}
	var roles model.Role
	var authentication model.Authentication
	db.Debug().Where("user_id = ?", userid).Find(&authentication)
	db.Debug().Where("ID =?", authentication.RoleID).Find(&roles)
	checkLogin := CheckLogins{
		Role:   roles.Name,
		Status: "ok",
	}
	b, _ := json.Marshal(checkLogin)
	fmt.Fprint(w, string(b))
}
