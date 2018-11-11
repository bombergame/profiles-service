package grpc

import (
	"context"
	"github.com/bombergame/common/logs"
	"github.com/bombergame/profiles-service/config"
	"github.com/bombergame/profiles-service/repositories"
	"google.golang.org/grpc"
	"net"
)

type Service struct {
	config *Config
	server *grpc.Server
}

type Config struct {
	Logger     *logs.Logger
	Repository repositories.ProfileRepository
}

func NewService(c *Config) *Service {
	srv := &Service{
		config: c,
	}

	grpcSrv := grpc.NewServer()
	RegisterProfilesServiceServer(grpcSrv, srv)

	srv.server = grpcSrv

	return srv
}

func (srv *Service) Run() error {
	addr := ":" + config.GrpcPort

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	srv.config.Logger.Info("grpc service running on: " + addr)
	return srv.server.Serve(lis)
}

func (srv *Service) Shutdown() error {
	srv.config.Logger.Info("grpc service shutdown initialized")
	srv.server.GracefulStop()
	return nil
}

func (srv *Service) IncProfileScore(ctx context.Context, req *ProfileID) (*Void, error) {
	return &Void{}, nil //TODO
}

func (srv *Service) GetProfileIDByCredentials(ctx context.Context, req *Credentials) (*ProfileID, error) {
	id, err := srv.config.Repository.FindIDByCredentials(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	profileID := &ProfileID{
		Value: *id,
	}

	return profileID, nil
}
