package alert_mongo

import (
	"context"
	"log/slog"

	"github.com/WiaanB/sisyphus/internal/alerting/domain"
	"github.com/WiaanB/sisyphus/internal/logging"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoAlertRepository struct {
	db *mongo.Client
}

func NewMongoAlertRepository(db *mongo.Client) *MongoAlertRepository {
	return &MongoAlertRepository{db: db}
}

func (r *MongoAlertRepository) GetAlertByRule(
	ctx context.Context,
	ruleID string,
) (*domain.Alert, error) {
	coll := r.db.Database("sisyphus").Collection("alerts")

	var alert domain.Alert
	err := coll.FindOne(ctx, bson.D{{"rule_id", ruleID}}).Decode(&alert)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			logging.Logger.Service.Error("not found", slog.String("rule_id", ruleID))
			return nil, nil
		}

		logging.Logger.Service.Error("failed to find the alert", slog.String("rule_id", ruleID))
		return nil, err
	}

	return &alert, nil
}

func (r *MongoAlertRepository) SaveAlert(
	ctx context.Context,
	alert *domain.Alert,
) error {
	coll := r.db.Database("sisyphus").Collection("alerts")

	opts := options.Replace().SetUpsert(true)
	filter := bson.D{{"id", alert.ID}}

	_, err := coll.ReplaceOne(ctx, filter, alert, opts)
	if err != nil {
		logging.Logger.Service.Error("failed to update alert", slog.String("err", err.Error()))
		return err
	}

	return err
}
