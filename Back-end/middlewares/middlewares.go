package middlewares

import (
	"context"
	"fmt"
	"hotel/auth"
	"net/http"
)

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Token := auth.TokenValid(r)
		if Token == nil {
			fmt.Fprint(w, "Phải đăng nhập trước khi đến trang web này")
			fmt.Fprint(w, "Render tới trang web")
			return
		} else {
			ctx := context.WithValue(r.Context(), "user_id", Token["user_id"])
			r = r.WithContext(ctx)
		}
		next(w, r)

		// r.Header.Set("user_id", "1")
	}
}
