package main

import (
	"github.com/bombergame/common/logs"
	"github.com/bombergame/profiles-service/clients/auth-service/grpc"
	"github.com/bombergame/profiles-service/repositories/postgres"
	"github.com/bombergame/profiles-service/services/grpc"
	"github.com/bombergame/profiles-service/services/rest"
	"os"
	"os/signal"
)

func main() {
	logger := logs.NewLogger()

	conn := postgres.NewConnection()

	defer conn.Close()
	if err := conn.Open(); err != nil {
		logger.Fatal(err)
		return
	}

	profileRepository := postgres.NewProfileRepository(conn)

	authGrpc := authgrpc.NewClient(
		&authgrpc.Config{
			Logger: logger,
		},
	)

	defer authGrpc.Disconnect()
	if err := authGrpc.Connect(); err != nil {
		logger.Fatal(err)
		return
	}

	restSrv := rest.NewService(
		&rest.ServiceConfig{
			Logger:            logger,
			AuthGrpc:          authGrpc,
			ProfileRepository: profileRepository,
		},
	)

	grpcSrv := profilesgrpc.NewService(
		&profilesgrpc.Config{
			Logger:     logger,
			Repository: profileRepository,
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
