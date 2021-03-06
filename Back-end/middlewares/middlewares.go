package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"hotel/auth"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		Token := auth.TokenValid(r)
		if Token == nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Ném mã xác thực vào!")
			return
		} else {
			v := jwt.MapClaims{
				"user_id":  Token["user_id"],
				"roles_id": Token["roles_id"],
			}
			data := Pretty(v)
			ctxUserId := context.WithValue(r.Context(), "data", data)
			r = r.WithContext(ctxUserId)
			next(w, r)
		}
		// r.Header.Set("user_id", "1")
	}
}
func SetMiddlewareAuthenticationUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		Token := auth.TokenValid(r)
		if Token == nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Ném mã xác thực vào!")
			return
		} else {
			v := jwt.MapClaims{
				"user_id":  Token["user_id"],
				"roles_id": Token["roles_id"],
			}
			data := Pretty(v)
			ctxUserId := context.WithValue(r.Context(), "data", data)
			r = r.WithContext(ctxUserId)
			if ConvertDataToken(data, "roles_id") != "1" {
				fmt.Fprintln(w, "Đây phải là tài khoản của User!")
			} else {
				next(w, r)
			}
		}
		// r.Header.Set("user_id", "1")
	}
}
func SetMiddlewareAuthenticationHotelOwner(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		Token := auth.TokenValid(r)
		if Token == nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Ném mã xác thực vào!")
			return
		} else {
			v := jwt.MapClaims{
				"user_id":  Token["user_id"],
				"roles_id": Token["roles_id"],
			}
			data := Pretty(v)
			ctxUserId := context.WithValue(r.Context(), "data", data)
			r = r.WithContext(ctxUserId)
			if ConvertDataToken(data, "roles_id") != "2" {
				fmt.Fprintln(w, "Đây phải là tài khoản của HotelOwner!")
			} else {
				next(w, r)
			}
		}
		// r.Header.Set("user_id", "1")
	}
}
func Pretty(data interface{}) string {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(b)
	// b1 := strings.Split(string(b), ",")
	// b2 := strings.Split(string(b1[0]), ": ")
	// b3 := strings.Split(string(b1[1]), ": ")
	// b4 := strings.Split(string(b3[1]), "\n")
	// i, _ := strconv.ParseUint(b2[1], 10, 64)
	// i1, _ := strconv.ParseUint(b4[0], 10, 64)
	// arr[0] = i
	// arr[1] = i1
	// return arr
}
func ConvertDataToken(DataToken interface{}, data string) string {
	str := fmt.Sprintf("%v", DataToken)
	sec := map[string]interface{}{}
	err := json.Unmarshal([]byte(str), &sec)
	if err != nil {
		fmt.Print(err)
	}
	str1 := fmt.Sprintf("%v", sec[data])
	return str1
}
