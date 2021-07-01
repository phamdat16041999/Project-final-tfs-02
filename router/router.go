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
