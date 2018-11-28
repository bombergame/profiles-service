package profilesrest

import (
	"github.com/bombergame/auth-service/services/grpc"
	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/rest"
	"github.com/bombergame/profiles-service/config"
	"github.com/bombergame/profiles-service/repositories"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type Service struct {
	rest.Service
	config     ServiceConfig
	components ServiceComponents
}

type ServiceConfig struct {
	rest.Config
}

type ServiceComponents struct {
	rest.Components
	ProfileRepository repositories.ProfileRepository
	AuthClient        *authgrpc.Client
}

func NewService(cf ServiceConfig, cp ServiceComponents) *Service {
	cf.Host, cf.Port = consts.EmptyString, config.HttpPort

	srv := &Service{
		config:     cf,
		components: cp,
		Service: *rest.NewService(
			cf.Config, cp.Components,
		),
	}

	mx := mux.NewRouter()

	mx.Handle("/", handlers.MethodHandler{
		http.MethodGet:  http.HandlerFunc(srv.getProfiles),
		http.MethodPost: http.HandlerFunc(srv.createProfile),
	})

	mx.Handle("/{profile_id:[0-9]+}", handlers.MethodHandler{
		http.MethodGet:    http.HandlerFunc(srv.getProfile),
		http.MethodPatch:  srv.WithAuth(http.HandlerFunc(srv.updateProfile)),
		http.MethodDelete: srv.WithAuth(http.HandlerFunc(srv.deleteProfile)),
	})

	srv.SetHandler(srv.WithLogs(srv.WithRecover(mx)))

	return srv
}
