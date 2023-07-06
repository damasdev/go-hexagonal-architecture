package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoDB *mongo.Database
)

type Mongo struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	UseSRV   string
}

func LoadMongoConfig() Mongo {
	conf := Mongo{
		Driver:   os.Getenv("MONGO_DRIVER"),
		Host:     os.Getenv("MONGO_HOST"),
		Port:     os.Getenv("MONGO_PORT"),
		User:     os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
		DBName:   os.Getenv("MONGO_DBNAME"),
		UseSRV:   os.Getenv("MONGO_SRV"),
	}
	return conf
}

func ConnectMongoDB() {
	conf := LoadMongoConfig()
	useSRV, _ := strconv.ParseBool(conf.UseSRV)
	connectionString := fmt.Sprintf(`mongodb+srv://%s:%s@%s`, conf.User, conf.Password, conf.Host)
	if !useSRV {
		connectionString = fmt.Sprintf(`mongodb://%s:%s`, conf.Host, conf.Port)
		if conf.User != "" {
			connectionString = fmt.Sprintf(`mongodb://%s:%s@%s:%s`, conf.User, conf.Password, conf.Host, conf.Port)
		}
	}

	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	MongoDB = mongoClient.Database(conf.DBName)
}
