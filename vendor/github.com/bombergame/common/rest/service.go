package rest

import (
	"context"
	"github.com/bombergame/common/auth"
	"github.com/bombergame/common/logs"
	"net/http"
)

type Service struct {
	config     Config
	components Components
	server     http.Server
}

type Config struct {
	Host string
	Port string
}

type Components struct {
	Logger      *logs.Logger
	AuthManager auth.AuthenticationManager
}

func NewService(config Config, components Components) *Service {
	return &Service{
		config:     config,
		components: components,
		server: http.Server{
			Addr: config.Host + ":" + config.Port,
		},
	}
}

func (srv *Service) Run() error {
	srv.components.Logger.Info("http service running on: " + srv.server.Addr)
	return srv.server.ListenAndServe()
}

func (srv *Service) Shutdown() error {
	srv.components.Logger.Info("http service shutdown initialized")
	return srv.server.Shutdown(context.TODO())
}

func (srv *Service) SetHandler(h http.Handler) {
	srv.server.Handler = h
}

func (srv *Service) Logger() *logs.Logger {
	return srv.components.Logger
}
