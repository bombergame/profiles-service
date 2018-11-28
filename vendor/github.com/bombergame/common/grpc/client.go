package grpc

import (
	"github.com/bombergame/common/logs"
	"google.golang.org/grpc"
)

type Client struct {
	Config     ClientConfig
	Components ClientComponents
	Conn       *grpc.ClientConn
}

type ClientConfig struct {
	ServiceHost string
	ServicePort string
}

type ClientComponents struct {
	Logger *logs.Logger
}

func NewClient(cf ClientConfig, cp ClientComponents) *Client {
	return &Client{
		Config:     cf,
		Components: cp,
	}
}

func (c *Client) Connect() error {
	var err error
	c.Conn, err = grpc.Dial(
		c.Config.ServiceHost+":"+c.Config.ServicePort,
		grpc.WithInsecure(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Disconnect() error {
	return c.Conn.Close()
}

func (c *Client) Logger() *logs.Logger {
	return c.Components.Logger
}
