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
	post.Path("/createbill").HandlerFunc(middlewares.SetMiddlewareAuthentication(model.Createbill))
	post.Path("/createhotel").HandlerFunc(middlewares.SetMiddlewareAuthentication(model.CreateHotel))

	//get method
	get := r.Methods(http.MethodGet).Subrouter()
	get.Path("/homepage").HandlerFunc(model.DataHomePage)
<<<<<<< HEAD
	get.Path("/hotel/{address}").HandlerFunc(model.GetHotelAddress)
=======
	get.Path("/homepage/hotel/{address}/{rate}").HandlerFunc(model.GetHotelAddress)
>>>>>>> d6d66811dd3574ee337fb74d350d3d08f074d059
	get.Path("/tophotel").HandlerFunc(model.GetTopHotel)
	get.Path("/detailhotel/{id}").HandlerFunc(model.GetDetailHotel)

	// methodput
	r.HandleFunc("/update/{id}", middlewares.SetMiddlewareAuthentication(model.UpdateAccount)).Methods("PUT")
	r.HandleFunc("/test", middlewares.SetMiddlewareAuthentication(Test)).Methods("PUT")
	http.Handle("/", r)
	//methoddelete
	r.HandleFunc("/delete/{id}", middlewares.SetMiddlewareAuthentication(model.DeleteAccount)).Methods("Delete")

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS", "PUT"},
	}).Handler(r)
	http.ListenAndServe(":8000", handler)
}
func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "test")
}
