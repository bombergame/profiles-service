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

	StorageUser     = env.GetVar("SESSIONS_STORAGE_USER", "user")
	StoragePassword = env.GetVar("SESSIONS_STORAGE_PASSWORD", "password")
	StorageHost     = env.GetVar("SESSIONS_STORAGE_HOST", "127.0.0.1")
	StoragePort     = env.GetVar("SESSIONS_STORAGE_PORT", "3306")
	StorageName     = env.GetVar("SESSIONS_STORAGE_NAME", "sessions")

	ProfilesServiceGrpcHost = env.GetVar("PROFILES_SERVICE_GRPC_HOST", "profiles-service")
	ProfilesServiceGrpcPort = env.GetVar("PROFILES_SERVICE_GRPC_PORT", "3000")
)
