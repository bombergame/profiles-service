package rest

import (
	"github.com/bombergame/common/logs"
	"github.com/bombergame/profiles-service/config"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type Service struct {
	server http.Server
	logger *logs.Logger
}

func NewService() *Service {
	srv := &Service{
		server: http.Server{
			Addr: ":" + config.HttpPort,
		},
		logger: logs.NewLogger(),
	}

	mx := mux.NewRouter()

	mx.Handle("/", handlers.MethodHandler{
		http.MethodGet:  http.HandlerFunc(srv.getProfiles),
		http.MethodPost: http.HandlerFunc(srv.createProfile),
	})

	mx.Handle("/{profile_id:[0-9]+}", handlers.MethodHandler{
		http.MethodGet:    http.HandlerFunc(srv.getProfile),
		http.MethodPatch:  http.HandlerFunc(srv.updateProfile),
		http.MethodDelete: http.HandlerFunc(srv.deleteProfile),
	})

	srv.server.Handler = srv.withLogs(srv.withRecover(mx))

	return srv
}

func (srv *Service) Run() error {
	return srv.server.ListenAndServe()
}
