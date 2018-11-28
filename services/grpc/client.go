package profilesgrpc

import (
	"github.com/bombergame/common/errs"
	"github.com/bombergame/common/grpc"
	"strings"
)

type Client struct {
	grpc.Client
	config     ClientConfig
	components ClientComponents
	client     ProfilesServiceClient
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

	c.client = NewProfilesServiceClient(c.Conn)

	addr := c.Config.ServiceHost + ":" + c.Config.ServicePort
	c.Logger().Info("profiles-service grpc client connection: " + addr)

	return nil
}

func (c *Client) Disconnect() error {
	return c.Client.Disconnect()
}

func (c *Client) IncProfileScore(profileID *ProfileID) (*Void, error) {
	return nil, nil //TODO
}

func (c *Client) GetProfileIDByCredentials(in *Credentials) (*ProfileID, error) {
	return nil, nil //TODO
}

func (c *Client) wrapError(err error) error {
	errMsg := err.Error()
	if strings.Contains(errMsg, "not found") {
		return errs.NewNotFoundError("profile not found")
	}
	return errs.NewServiceError(err)
}
