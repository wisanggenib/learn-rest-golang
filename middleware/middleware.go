package middleware

import (
	"learn/REST/utils"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		if !ok {
			utils.ResponseJSON(w, "Not Authorized", http.StatusUnauthorized)
			return
		}
		//Just Simple Check, using basic auth if Username or password null got unauthorized
		if username == "" || password == "" {
			utils.ResponseJSON(w, "Please Fill Auth", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})

}
