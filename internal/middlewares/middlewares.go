package middlewares

import (
	"log"
	"net/http"
)

func BasicAuth(isUserAuth func(name, pwd string) error) func(next http.Handler) http.Handler {
	fn := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, pass, ok := r.BasicAuth()
			if !ok {
				log.Printf("Auth Failed")
				basicAuthFailed(w)
				return
			}

			err := isUserAuth(user, pass)
			if err != nil {
				log.Printf("Auth Failed")
				basicAuthFailed(w)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	return fn
}

func basicAuthFailed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
}
