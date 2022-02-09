package mongorm

import (
	"context"
	"time"
)

type Config struct {
	dbName            string
	host              string
	port              string
	uri               string
	contextFn         func() context.Context
	connectionTimeOut time.Duration
}

func NewConfig() *Config {
	return &Config{
		contextFn:         context.Background,
		connectionTimeOut: time.Second * 10,
	}
}

func (c *Config) SetDBName(name string) *Config {
	c.dbName = name
	return c
}

func (c *Config) SetHost(host string) *Config {
	c.uri = ""
	c.host = host
	return c
}

func (c *Config) SetPort(port string) *Config {
	c.uri = ""
	c.port = port
	return c
}

func (c *Config) SetURI(uri string) *Config {
	c.host, c.port = "", ""
	c.uri = uri
	return c
}

func (c *Config) SetContextFn(cfn func() context.Context) *Config {
	c.contextFn = cfn
	return c
}

func (c *Config) SetConnectionTimeOut(dur time.Duration) *Config {
	c.connectionTimeOut = dur
	return c
}
