package mongorm

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	defaultInstance *Mongorm
	defaultDatabase *mongo.Database
)

type Mongorm struct {
	client *mongo.Client
	config *Config
}

func New(config *Config) (m *Mongorm, err error) {
	if config == nil {
		config = NewConfig()
	}
	if config.contextFn == nil {
		config.contextFn = context.Background
	}
	ctx, cancel := context.WithTimeout(config.contextFn(), config.connectionTimeOut)
	defer cancel()

	var co *options.ClientOptions
	co = options.Client()
	if config.uri != "" {
		co.ApplyURI(config.uri)
	} else if config.host != "" && config.port != "" {
		co.ApplyURI(fmt.Sprintf("mongodb://%s:%s", config.host, config.port))
	}

	var client *mongo.Client
	client, err = mongo.Connect(ctx, co)
	if err != nil {
		return
	}

	m = &Mongorm{client: client, config: config}
	defaultInstance = m
	defaultDatabase = m.client.Database(config.dbName)
	return
}

func (m *Mongorm) Close() error {
	return m.client.Disconnect(m.config.contextFn())
}
