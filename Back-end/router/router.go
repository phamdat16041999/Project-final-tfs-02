package router

import (
	"fmt"
	"hotel/middlewares"
	"hotel/model"

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
	post.Path("/changepassword").HandlerFunc(middlewares.SetMiddlewareAuthenticationUser(model.ChangePassword))
	post.Path("/active").HandlerFunc(model.ActiveAccount)
	post.Path("/createbill").HandlerFunc(middlewares.SetMiddlewareAuthenticationUser(model.Createbill))
	post.Path("/rating").HandlerFunc(middlewares.SetMiddlewareAuthenticationUser(model.Rating))
	post.Path("/checkroomstatus").HandlerFunc(middlewares.SetMiddlewareAuthenticationUser(model.Checkroomstatus))
	post.Path("/createhotel").HandlerFunc(middlewares.SetMiddlewareAuthenticationUser(model.CreateHotel))
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
	get.Path("/hotelier").HandlerFunc(middlewares.SetMiddlewareAuthenticationUser(model.Hotelier))
	get.Path("/option").HandlerFunc(model.OptionforHotel)

	get.Path("/detailbillofmanagerhotel").HandlerFunc(middlewares.SetMiddlewareAuthenticationUser(model.Allbillofmanagerhotel))
	get.Path("/detailbill/{id}").HandlerFunc(middlewares.SetMiddlewareAuthenticationUser(model.Detailbill))
	// delete method
	detelte := r.Methods(http.MethodDelete).Subrouter()
	detelte.Path("/hotel/{id}").HandlerFunc(middlewares.SetMiddlewareAuthenticationUser(model.DeleteHotel))
	// methodput
	r.HandleFunc("/update/{id}", middlewares.SetMiddlewareAuthenticationUser(model.UpdateAccount)).Methods("PUT")
	r.HandleFunc("/checklogin", middlewares.SetMiddlewareAuthenticationUser(CheckLogin)).Methods("GET")
	http.Handle("/", r)
	//methoddelete
	r.HandleFunc("/delete/{id}", middlewares.SetMiddlewareAuthenticationUser(model.DeleteAccount)).Methods("Delete")
	// chat API
	r.HandleFunc("/ws", websocket.HandleConnections)

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS", "PUT"},
		AllowedHeaders: []string{"*"},
	}).Handler(r)
	http.ListenAndServe(":8080", handler)
}
func CheckLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
