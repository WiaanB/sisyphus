package main

import (
	"context"
	"os"

	"github.com/WiaanB/sisyphus/internal/alerting"
	"github.com/WiaanB/sisyphus/internal/alerting/domain"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Client *mongo.Client

func main() {
	// setup
	loadEnv()
	ConnectToMongo()
	defer func() {
		if err := Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// TODO: Add logger fmt.Println("Creating an Alert Service")
	all := alerting.New(Client)

	err := all.Service.EvaluateRule(context.TODO(), domain.Rule{
		ID:        "1",
		Metric:    "Name",
		Threshold: 10,
		Severity:  "critical",
	}, 11)
	if err != nil {
		panic(err)
	}
}

func loadEnv() {
	_ = godotenv.Load()
}

func ConnectToMongo() {
	var uri string
	if uri = os.Getenv("MONGODB_URI"); uri == "" {
		// TODO: Add logger log.Fatal("You must set your 'MONGODB_URI' environment variable.")
	}
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	var err error
	Client, err = mongo.Connect(opts)

	if err != nil {
		panic(err)
	}

	var result bson.M
	if err = Client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	// TODO: Add logger fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
