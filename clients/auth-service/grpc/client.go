package grpc

import (
	"context"
	"github.com/bombergame/common/errs"
	"github.com/bombergame/common/logs"
	"github.com/bombergame/profiles-service/config"
	"google.golang.org/grpc"
	"strings"
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
	id, err := c.client.GetProfileID(context.TODO(), authInfo)
	if err != nil {
		return nil, c.wrapError(err)
	}
	return id, nil
}

func (c *Client) DeleteAllSessions(profileID *ProfileID) error {
	_, err := c.client.DeleteAllSessions(context.TODO(), profileID)
	if err != nil {
		return c.wrapError(err)
	}
	return nil
}

func (c *Client) wrapError(err error) error {
	errMsg := err.Error()

	if strings.Contains(errMsg, errs.InvalidFormatErrorMessagePrefix) {
		return errs.NewNotAuthorizedError()
	}
	if strings.Contains(errMsg, errs.NotAuthorizedErrorMessage) {
		return errs.NewNotAuthorizedError()
	}

	return errs.NewServiceError(err)
}
