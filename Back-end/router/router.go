package router

import (
	"fmt"
	"hotel/middlewares"
	"hotel/model"
	"hotel/pkg/websocket"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Run() {
	r := mux.NewRouter().StrictSlash(true)
	post := r.Methods(http.MethodPost).Subrouter()
	post.Path("/account").HandlerFunc(model.CreateAccount)
	post.Path("/login").HandlerFunc(model.LoginAcount)
	post.Path("/forgotpassword").HandlerFunc(model.ForgotPassword)
	post.Path("/changepassword").HandlerFunc(middlewares.SetMiddlewareAuthentication(model.ChangePassword))
	post.Path("/active").HandlerFunc(middlewares.SetMiddlewareAuthentication(model.ActiveAccount))
	post.Path("/createbill").HandlerFunc(middlewares.SetMiddlewareAuthentication(model.Createbill))
	post.Path("/rating").HandlerFunc(middlewares.SetMiddlewareAuthentication(model.Rating))
	post.Path("/checkroomstatus").HandlerFunc(middlewares.SetMiddlewareAuthentication(model.Checkroomstatus))
	// post.Path("/payment").HandlerFunc(middlewares.SetMiddlewareAuthentication(model.Payment))

	//get method
	get := r.Methods(http.MethodGet).Subrouter()
	get.Path("/homepage").HandlerFunc(model.DataHomePage)
	get.Path("/hotel/{address}/{rate}").HandlerFunc(model.SearchHotelAddress)
	get.Path("/tophotel").HandlerFunc(model.GetTopHotel)
	get.Path("/detailhotel/{id}").HandlerFunc(model.GetDetailHotel)
	get.Path("/search/{name}").HandlerFunc(model.SearchByName)

	// methodput
	r.HandleFunc("/update/{id}", middlewares.SetMiddlewareAuthentication(model.UpdateAccount)).Methods("PUT")
	r.HandleFunc("/CheckLogin", middlewares.SetMiddlewareAuthentication(CheckLogin)).Methods("GET")
	r.HandleFunc("/test", middlewares.SetMiddlewareAuthentication(Test)).Methods("PUT")
	http.Handle("/", r)
	//methoddelete
	r.HandleFunc("/delete/{id}", middlewares.SetMiddlewareAuthentication(model.DeleteAccount)).Methods("Delete")
	// chat API
	r.HandleFunc("/ws", websocket.HandleConnections)

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS", "PUT"},
	}).Handler(r)
	http.ListenAndServe(":8080", handler)
}
func CheckLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "test")
}
