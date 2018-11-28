package config

import (
	"github.com/bombergame/common/args"
	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/env"
)

var (
	HttpPort = args.GetString("http_port", "80")
	GrpcPort = args.GetString("grpc_port", "3000")

	TokenSignKey = env.GetVar("TOKEN_SIGN_KEY", consts.EmptyString)

	StorageHost     = env.GetVar("PROFILES_STORAGE_HOST", "127.0.0.1")
	StoragePort     = env.GetVar("PROFILES_STORAGE_PORT", "5432")
	StorageName     = env.GetVar("PROFILES_STORAGE_NAME", "profiles")
	StorageUser     = env.GetVar("PROFILES_STORAGE_USER", "user")
	StoragePassword = env.GetVar("PROFILES_STORAGE_PASSWORD", "password")
	StorageSSLMode  = env.GetVar("PROFILES_STORAGE_SSL_MODE", "disable")

	AuthServiceGrpcHost = env.GetVar("AUTH_SERVICE_GRPC_HOST", "auth-service")
	AuthServiceGrpcPort = env.GetVar("AUTH_SERVICE_GRPC_PORT", "3000")
)
