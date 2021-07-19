package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"hotel/auth"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
			// data := Pretty(r.Context().Value("data"))
			// fmt.Fprintln(w, data)
		}
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
		// r.Header.Set("user_id", "1")
	}
}
func Pretty(data interface{}) [2]uint64 {
	var arr [2]uint64
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
	}
	b1 := strings.Split(string(b), ",")
	b2 := strings.Split(string(b1[0]), ": ")
	b3 := strings.Split(string(b1[1]), ": ")
	b4 := strings.Split(string(b3[1]), "\n")
	i, _ := strconv.ParseUint(b2[1], 10, 64)
	i1, _ := strconv.ParseUint(b4[0], 10, 64)
	arr[0] = i
	arr[1] = i1
	return arr
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
