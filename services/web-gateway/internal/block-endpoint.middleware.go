package internal

import (
	"net/http"
)

func BlockEndpointMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request should be blocked
		if shouldBlockRequest(r) {
			http.Error(w, "Access to this endpoint is blocked", http.StatusForbidden)
			return
		}

		// Call the next handler if the request is allowed
		next.ServeHTTP(w, r)
	})
}

func shouldBlockRequest(r *http.Request) bool {
	//if r.URL.Path == "/v1/party" && r.Method == "PUT" {
	//	return true
	//}

	return false
}
