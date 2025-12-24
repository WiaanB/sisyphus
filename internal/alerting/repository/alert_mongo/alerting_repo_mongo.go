package alert_mongo

import (
	"context"

	"github.com/WiaanB/sisyphus/internal/alerting/domain"
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
	// TODO: Add logger

	curr, err := coll.Find(ctx, bson.D{{"ruleId", ruleID}})
	if err != nil {
		return nil, err
	}

	var alert domain.Alert
	err = curr.Decode(&alert)
	if err != nil {
		return nil, err
	}

	return &alert, nil
}

func (r *MongoAlertRepository) SaveAlert(
	ctx context.Context,
	alert *domain.Alert,
) error {
	coll := r.db.Database("sisyphus").Collection("alerts")
	// TODO: Add logger

	opts := options.Replace().SetUpsert(true)
	filter := bson.D{{"id", alert.ID}}

	_, err := coll.ReplaceOne(ctx, filter, alert, opts)

	return err
}
