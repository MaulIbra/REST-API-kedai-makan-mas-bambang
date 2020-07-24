package middleware

import (
	"fmt"
	"github.com/maulIbra/clean-architecture-go/utils"
	"net/http"
)

func TokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenValue := r.Header.Get("token")
		if len(tokenValue) == 0 {
			utils.HandleResponseError(w,http.StatusUnauthorized,"Unauthorized")
		}else{
			_, err := utils.JwtDecoder(tokenValue)
			if err != nil {
				fmt.Println(err)
				utils.HandleResponseError(w,http.StatusUnauthorized,"Expired Token")
			}else{
				next.ServeHTTP(w, r)
			}
		}
	})
}

