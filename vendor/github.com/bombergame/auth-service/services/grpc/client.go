package authgrpc

import (
	"context"
	"github.com/bombergame/common/errs"
	"github.com/bombergame/common/grpc"
)

type Client struct {
	grpc.Client
	config     ClientConfig
	components ClientComponents
	client     AuthServiceClient
}

type ClientConfig struct {
	grpc.ClientConfig
}

type ClientComponents struct {
	grpc.ClientComponents
}

func NewClient(cf ClientConfig, cp ClientComponents) *Client {
	return &Client{
		config:     cf,
		components: cp,
		Client: *grpc.NewClient(
			cf.ClientConfig,
			cp.ClientComponents,
		),
	}
}

func (c *Client) Connect() error {
	err := c.Client.Connect()
	if err != nil {
		return err
	}

	c.client = NewAuthServiceClient(c.Conn)

	addr := c.Config.ServiceHost + ":" + c.Config.ServicePort
	c.Logger().Info("auth-service grpc client connection: " + addr)

	return nil
}

func (c *Client) Disconnect() error {
	return c.Client.Disconnect()
}

func (c *Client) DeleteAllSessions(id ProfileID) error {
	_, err := c.client.DeleteAllSessions(context.TODO(), &id)
	if err != nil {
		return c.wrapError(err)
	}
	return nil
}

func (c *Client) wrapError(err error) error {
	return errs.NewServiceError(err)
}
