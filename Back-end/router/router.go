package router

import (
	"fmt"
	"hotel/middlewares"
	"hotel/model"
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
	post.Path("/changepassword").HandlerFunc(model.ChangePassword)
	post.Path("/active").HandlerFunc(model.ActiveAccount)

	//get method
	get := r.Methods(http.MethodGet).Subrouter()
	get.Path("/homepage").HandlerFunc(model.DataHomePage)
	get.Path("/homepage/hotel/{address}").HandlerFunc(model.GetHotelAddress)
	get.Path("/homepage/tophotel").HandlerFunc(model.TopHotel)
	get.Path("/homepage/eachhotel/{id}").HandlerFunc(model.GetEachHotel)

	// methodput
	r.HandleFunc("/test", middlewares.SetMiddlewareAuthentication(Test)).Methods("PUT")
	http.Handle("/", r)

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS", "PUT"},
	}).Handler(r)
	http.ListenAndServe(":8000", handler)
}
func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "test")
}
