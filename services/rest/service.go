package rest

import (
	"context"
	"github.com/bombergame/common/logs"
	authgrpc "github.com/bombergame/profiles-service/clients/auth-service/grpc"
	"github.com/bombergame/profiles-service/config"
	"github.com/bombergame/profiles-service/repositories"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type Service struct {
	config *Config
	server http.Server
}

type Config struct {
	Logger            *logs.Logger
	AuthGrpc          *authgrpc.Client
	ProfileRepository repositories.ProfileRepository
}

func NewService(c *Config) *Service {
	srv := &Service{
		config: c,
		server: http.Server{
			Addr: ":" + config.HttpPort,
		},
	}

	mx := mux.NewRouter()

	mx.Handle("/", handlers.MethodHandler{
		http.MethodGet:  http.HandlerFunc(srv.getProfiles),
		http.MethodPost: http.HandlerFunc(srv.createProfile),
	})

	mx.Handle("/{profile_id:[0-9]+}", handlers.MethodHandler{
		http.MethodGet:    http.HandlerFunc(srv.getProfile),
		http.MethodPatch:  srv.withAuthRestrict(http.HandlerFunc(srv.updateProfile)),
		http.MethodDelete: srv.withAuthRestrict(http.HandlerFunc(srv.deleteProfile)),
	})

	srv.server.Handler = srv.withLogs(srv.withRecover(mx))

	return srv
}

func (srv *Service) Run() error {
	srv.config.Logger.Info("http service running on: " + srv.server.Addr)
	return srv.server.ListenAndServe()
}

func (srv *Service) Shutdown() error {
	srv.config.Logger.Info("http service shutdown initialized")
	return srv.server.Shutdown(context.TODO())
}
