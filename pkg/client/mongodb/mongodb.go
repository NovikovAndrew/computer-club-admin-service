package mongodb

import (
	"context"
	"fmt"
	"net/url"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConfig struct {
	Host     string            `yaml:"host"`
	Port     string            `yaml:"port"`
	Username string            `yaml:"username"`
	Password string            `yaml:"password"`
	Database string            `yaml:"database"`
	AuthDB   string            `yaml:"authdb"`
	Options  map[string]string `yaml:"options"`
}

const (
	mongoDBAuthFormat = "mongodb://%s%s@%s:%s?%s"
	mongoDBFormat     = "mongodb://%s:%s?%s"
)

func MongoClient(ctx context.Context, cfg *MongoDBConfig) (db *mongo.Database, err error) {
	var isAuth bool
	var connectionURI string
	items := url.Values{}
	for k, v := range cfg.Options {
		items.Add(k, v)
	}

	if cfg.Username == "" && cfg.Password == "" {
		connectionURI = fmt.Sprintf(mongoDBFormat, cfg.Host, cfg.Port, items.Encode())
	} else {
		isAuth = true
		connectionURI = fmt.Sprintf(mongoDBAuthFormat, cfg.Username, cfg.Password, cfg.Host, cfg.Port, items.Encode())
	}

	clientOptins := options.Client().ApplyURI(connectionURI)

	if isAuth {
		if cfg.AuthDB == "" {
			cfg.AuthDB = cfg.Database
		}

		clientOptins.SetAuth(options.Credential{
			AuthSource: cfg.AuthDB,
			Username:   cfg.Username,
			Password:   cfg.Password,
		})
	}

	client, err := mongo.Connect(ctx, clientOptins)

	if err != nil {
		return nil, err
	}

	return client.Database(cfg.Database), err
}
