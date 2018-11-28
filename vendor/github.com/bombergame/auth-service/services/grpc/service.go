package authgrpc

import (
	"context"
	"github.com/bombergame/auth-service/config"
	"github.com/bombergame/auth-service/repositories"
	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/grpc"
)

type Service struct {
	grpc.Service
	config     ServiceConfig
	components ServiceComponents
}

type ServiceConfig struct {
	grpc.ServiceConfig
}

type ServiceComponents struct {
	grpc.ServiceComponents
	SessionRepository repositories.SessionRepository
}

func NewService(cf ServiceConfig, cp ServiceComponents) *Service {
	cf.Host, cf.Port = consts.EmptyString, config.GrpcPort

	srv := &Service{
		config:     cf,
		components: cp,
		Service: *grpc.NewService(
			cf.ServiceConfig,
			cp.ServiceComponents,
		),
	}

	RegisterAuthServiceServer(srv.Server, srv)

	return srv
}

func (srv *Service) DeleteAllSessions(ctx context.Context, id *ProfileID) (*Void, error) {
	err := srv.components.SessionRepository.DeleteAllSessions(id.Value)
	if err != nil {
		return nil, err
	}
	return &Void{}, err
}
