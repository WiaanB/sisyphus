package main

import (
	"context"
	"log/slog"

	"github.com/WiaanB/sisyphus/internal/alerting"
	"github.com/WiaanB/sisyphus/internal/alerting/domain"
	"github.com/WiaanB/sisyphus/internal/infra/db/mongo_db"
	"github.com/WiaanB/sisyphus/internal/logging"
	"github.com/joho/godotenv"
)

func main() {
	// setup
	loadEnv()
	logging.New()
	logging.Logger.Service.Info("logger started")

	mongo_db.ConnectToMongo()
	defer func() {
		if err := mongo_db.Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	all := alerting.New(mongo_db.Client)

	err := all.Service.EvaluateRule(context.TODO(), domain.Rule{
		ID:        "1",
		Metric:    "Name",
		Threshold: 10,
		Severity:  "critical",
	}, 9)
	if err != nil {
		logging.Logger.Service.Error("failed to evaluate rule", slog.String("err", err.Error()))
		panic(err)
	}
}

func loadEnv() {
	_ = godotenv.Load()
}
