package rest

import (
	"fmt"
	authgrpc "github.com/bombergame/profiles-service/clients/auth-service/grpc"
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
			userAgent, err := srv.readUserAgent(r)
			if err != nil {
				srv.writeErrorWithBody(w, err)
				return
			}

			authToken, err := srv.readAuthToken(r)
			if err != nil {
				srv.writeErrorWithBody(w, err)
				return
			}

			id, err := srv.config.AuthGrpc.GetProfileID(
				&authgrpc.AuthInfo{
					UserAgent: userAgent,
					Token:     authToken,
				},
			)
			if err != nil {
				srv.writeErrorWithBody(w, err)
				return
			}

			srv.setAuthProfileID(r, id.Value)
			h.ServeHTTP(w, r)
		},
	)
}
