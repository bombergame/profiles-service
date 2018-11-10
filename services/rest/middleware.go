package rest

import (
	"net/http"
)

func withRecover(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if r := recover(); r != nil {
					//TODO
				}
			}()
			h.ServeHTTP(w, r)
		},
	)
}

func withAuthPass(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			//TODO
			h.ServeHTTP(w, r)
		},
	)
}

func withAuthRestrict(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			//TODO
			h.ServeHTTP(w, r)
		},
	)
}
