package alerting

import (
	"github.com/WiaanB/sisyphus/internal/alerting/repository/alert_mongo"
	"github.com/WiaanB/sisyphus/internal/alerting/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Alerting struct {
	Service *service.AlertService
}

func New(mCl *mongo.Client) *Alerting {
	repo := alert_mongo.NewMongoAlertRepository(mCl)
	svc := service.NewAlertService(repo)

	return &Alerting{
		Service: svc,
	}
}
