package rest

import (
	"fmt"
	"net/http"
)

func (srv *Service) withRecover(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if r := recover(); r != nil {
					srv.config.Logger.Error(r)
					srv.writeError(w, http.StatusInternalServerError)
				}
			}()
			h.ServeHTTP(w, r)
		},
	)
}

func (srv *Service) withLogs(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			wr := &LoggingResponseWriter{
				writer: w,
			}
			h.ServeHTTP(wr, r)

			srv.config.Logger.Info(fmt.Sprintf("%s %s %d", r.Method, r.URL.Path, wr.status))
		},
	)
}

func (srv *Service) withAuthRestrict(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			//TODO
			h.ServeHTTP(w, r)
		},
	)
}
