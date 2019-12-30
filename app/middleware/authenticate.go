package middleware

import (
	"log"
	"net/http"
)

/**
 * Stop Request if there is a login
 * @param http.Handler  routes\AplicationV1Router next
 * @return http.Handler
 */
func LogginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.RequestURI, "middleware")
		next.ServeHTTP(writer, request)
	})
}
