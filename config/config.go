package config

import (
	"github.com/bombergame/common/args"
	"github.com/bombergame/common/env"
)

var (
	HttpPort = args.GetString("http_port", "80")
	GrpcPort = args.GetString("grpc_port", "3000")

	StorageHost     = env.GetVar("PROFILE_SERVICE_DB_HOST", "127.0.0.1")
	StoragePort     = env.GetVar("PROFILE_SERVICE_DB_PORT", "5432")
	StorageName     = env.GetVar("PROFILE_SERVICE_DB_NAME", "profiles")
	StorageUser     = env.GetVar("PROFILE_SERVICE_DB_USER", "user")
	StoragePassword = env.GetVar("PROFILE_SERVICE_DB_PASSWORD", "password")
	StorageSSLMode  = env.GetVar("PROFILE_SERVICE_DB_SSL_MODE", "disable")
)
