package middlewares

import (
	"fmt"
	"hotel/auth"
	"net/http"
)

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			fmt.Fprint(w, "Phải đăng nhập trước khi đến trang web này")
			fmt.Fprint(w, "Render tới trang web")
			return
		}
		next(w, r)
	}
}
