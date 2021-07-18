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
		Token := auth.TokenValid(r)
		if Token == nil {
			fmt.Fprint(w, "401")
			return
		} else {
			v := jwt.MapClaims{
				"user_id":  Token["user_id"],
				"roles_id": Token["roles_id"],
			}
			data := Pretty(v)
			ctxUserId := context.WithValue(r.Context(), "data", data)
			r = r.WithContext(ctxUserId)
			// data := Pretty(r.Context().Value("data"))
			// fmt.Fprintln(w, data)
		}
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
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
}
func ConvertDataToken(DataToken interface{}, data string) string {
	str := fmt.Sprintf("%v", DataToken)
	sec := map[string]interface{}{}
	if err := json.Unmarshal([]byte(str), &sec); err != nil {
		fmt.Print(err)
	}
	str1 := fmt.Sprintf("%v", sec[data])
	return str1
}
