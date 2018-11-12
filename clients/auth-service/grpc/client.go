package grpc

import (
	"github.com/bombergame/common/logs"
	"github.com/bombergame/profiles-service/config"
	"google.golang.org/grpc"
)

type Client struct {
	config *Config
	conn   *grpc.ClientConn
	client AuthServiceClient
}

type Config struct {
	Logger *logs.Logger
}

func NewClient(c *Config) *Client {
	return &Client{
		config: c,
	}
}

func (c *Client) Connect() error {
	var err error
	c.conn, err = grpc.Dial(config.AuthServiceGrpcAddress, grpc.WithInsecure())
	if err != nil {
		return c.wrapError(err)
	}

	c.client = NewAuthServiceClient(c.conn)

	c.config.Logger.Info("auth-service grpc connection: " + config.AuthServiceGrpcAddress)
	return nil
}

func (c *Client) Disconnect() error {
	c.config.Logger.Info("auth-service grpc connection shutdown")
	return c.conn.Close()
}

func (c *Client) GetProfileID(authInfo *AuthInfo) (*ProfileID, error) {
	return nil, nil //TODO
}

func (c *Client) DeleteAllSessions(profileID *ProfileID) error {
	return nil //TODO
}

func (c *Client) wrapError(err error) error {
	return nil //TODO
}
