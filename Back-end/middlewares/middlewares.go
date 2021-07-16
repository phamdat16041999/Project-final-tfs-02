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
			ctxUserId := context.WithValue(r.Context(), "data", v)
			r = r.WithContext(ctxUserId)
			data := Pretty(r.Context().Value("data"))
			fmt.Fprintln(w, data)
			// 	return data
		}
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
