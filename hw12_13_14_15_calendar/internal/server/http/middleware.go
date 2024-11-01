package internalhttp

import (
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler { //nolint:unused
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Host)
		log.Println(r.RemoteAddr)
		log.Println(r.Method)
		log.Println(r.UserAgent())

		next.ServeHTTP(w, r)
		//log.Println(r.Response.Status)
		//log.Println(r.Response.Body)
	})
}
