package grpc

import (
	"github.com/bombergame/common/auth"
	"github.com/bombergame/common/logs"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type Service struct {
	Config     ServiceConfig
	Components ServiceComponents
	Server     *grpc.Server
}

type ServiceConfig struct {
	Host string
	Port string
}

type ServiceComponents struct {
	Logger      *logs.Logger
	AuthManager auth.AuthenticationManager
}

func NewService(cf ServiceConfig, cp ServiceComponents) *Service {
	logEntry := logrus.NewEntry(cp.Logger.AsLogrusLogger())
	grpc_logrus.ReplaceGrpcLogger(logEntry)

	srv := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_recovery.UnaryServerInterceptor(),
			grpc_logrus.UnaryServerInterceptor(logEntry),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_recovery.StreamServerInterceptor(),
			grpc_logrus.StreamServerInterceptor(logEntry),
		),
	)

	return &Service{
		Config:     cf,
		Components: cp,
		Server:     srv,
	}
}

func (srv *Service) Run() error {
	addr := srv.Config.Host + ":" + srv.Config.Port

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	srv.Logger().Info("grpc service running on: " + addr)
	return srv.Server.Serve(lis)
}

func (srv *Service) Shutdown() error {
	srv.Logger().Info("grpc service shutdown initialized")
	srv.Server.GracefulStop()
	return nil
}

func (srv *Service) Logger() *logs.Logger {
	return srv.Components.Logger
}
