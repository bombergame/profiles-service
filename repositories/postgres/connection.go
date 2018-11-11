package postgres

import (
	"database/sql"
	"fmt"
	"github.com/bombergame/profiles-service/config"
	_ "github.com/lib/pq"
)

type Connection struct {
	str string
	db  *sql.DB
}

func NewConnection() *Connection {
	return &Connection{
		str: fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=%s",
			config.StorageUser, config.StoragePassword,
			config.StorageHost, config.StoragePort,
			config.StorageName, config.StorageSSLMode,
		),
	}
}

func (c *Connection) Open() error {
	var err error

	c.db, err = sql.Open("postgres", c.str)
	if err != nil {
		return err
	}

	return nil
}

func (c *Connection) Close() error {
	return c.db.Close()
}
