package mongo_db

import (
	"context"
	"log/slog"
	"os"

	"github.com/WiaanB/sisyphus/internal/logging"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Client *mongo.Client

func ConnectToMongo() {
	logging.Logger.Service.Info("connecting to mongodb")

	var uri string
	if uri = os.Getenv("MONGODB_URI"); uri == "" {
		logging.Logger.Service.Error("you must set your 'MONGODB_URI' environment variable.")
	}
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	var err error
	Client, err = mongo.Connect(opts)

	if err != nil {
		logging.Logger.Service.Error("failed to connect to MongoDB", slog.String("error", err.Error()))
	}

	var result bson.M
	if err = Client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		logging.Logger.Service.Error("failed to ping the MongoDB Instance", slog.String("error", err.Error()))
		panic(err)
	}
	logging.Logger.Service.Info("successfully connected to MongoDB")
}
