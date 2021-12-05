package mongodb

import (
	"context"
	"errors"
	"fmt"
	config "loanapp/config/env"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mutex                         = &sync.Mutex{}
	mongoClient   *mongo.Client   = nil
	mongoDatabase *mongo.Database = nil
)

func GetClient() (*mongo.Client, error) {
	if mongoClient == nil {
		return nil, errors.New("please init first")
	}

	return mongoClient, nil
}

func GetDatabase() (*mongo.Database, error) {
	if mongoClient == nil {
		return nil, errors.New("please init first")
	}

	return mongoDatabase, nil
}

func Init(envConfig *config.AppConfig) (*mongo.Client, error) {
	if mongoClient != nil {
		return mongoClient, nil
	}

	mutex.Lock()
	defer mutex.Unlock()

	//re-check after locking
	if mongoClient != nil {
		return mongoClient, nil
	}

	var prefix = "mongodb://"

	address := fmt.Sprintf("%s:%s", envConfig.MongoDB.DBAddress, envConfig.MongoDB.DBPort)

	// Set client options
	clientOptions := options.Client().ApplyURI(prefix + address)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	mongoClient = client
	mongoDatabase = client.Database(config.GetConfig().MongoDB.DBName)

	return mongoClient, nil
}
