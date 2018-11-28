package main

import (
	"github.com/bombergame/auth-service/services/grpc"
	"github.com/bombergame/common/grpc"
	"github.com/bombergame/common/logs"
	"github.com/bombergame/common/rest"
	"github.com/bombergame/profiles-service/auth"
	"github.com/bombergame/profiles-service/config"
	"github.com/bombergame/profiles-service/repositories/postgres"
	"github.com/bombergame/profiles-service/services/grpc"
	"github.com/bombergame/profiles-service/services/rest"
	"os"
	"os/signal"
)

func main() {
	logger := logs.NewLogger()

	conn := postgres.NewConnection()
	if err := conn.Open(); err != nil {
		logger.Fatal(err)
		return
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	profileRepository := postgres.NewProfileRepository(conn)

	authManager := auth.NewJwtAuthManager()

	authClient := authgrpc.NewClient(
		authgrpc.ClientConfig{
			ClientConfig: grpc.ClientConfig{
				ServiceHost: config.AuthServiceGrpcHost,
				ServicePort: config.AuthServiceGrpcPort,
			},
		},
		authgrpc.ClientComponents{
			ClientComponents: grpc.ClientComponents{
				Logger: logger,
			},
		},
	)

	if err := authClient.Connect(); err != nil {
		logger.Fatal(err)
		return
	}
	defer func() {
		err := authClient.Disconnect()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	restSrv := profilesrest.NewService(
		profilesrest.ServiceConfig{
			Config: rest.Config{},
		},
		profilesrest.ServiceComponents{
			Components: rest.Components{
				Logger:      logger,
				AuthManager: authManager,
			},
			ProfileRepository: profileRepository,
			AuthClient:        authClient,
		},
	)

	grpcSrv := profilesgrpc.NewService(
		profilesgrpc.ServiceConfig{
			ServiceConfig: grpc.ServiceConfig{},
		},
		profilesgrpc.ServiceComponents{
			ServiceComponents: grpc.ServiceComponents{
				Logger: logger,
			},
		},
	)

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)

	go func() {
		if err := restSrv.Run(); err != nil {
			logger.Fatal(err)
		}
	}()

	go func() {
		if err := grpcSrv.Run(); err != nil {
			logger.Fatal(err)
		}
	}()

	<-ch

	if err := restSrv.Shutdown(); err != nil {
		logger.Fatal(err)
	}

	if err := grpcSrv.Shutdown(); err != nil {
		logger.Fatal(err)
	}
}
