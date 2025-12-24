package service

import (
	"context"
	"time"

	"github.com/WiaanB/sisyphus/internal/alerting/domain"
	"github.com/google/uuid"
)

type AlertRepository interface {
	GetAlertByRule(ctx context.Context, ruleID string) (*domain.Alert, error)
	SaveAlert(ctx context.Context, alert *domain.Alert) error
}
type AlertService struct {
	repo  AlertRepository
	clock func() time.Time
}

func NewAlertService(repo AlertRepository) *AlertService {
	return &AlertService{
		repo:  repo,
		clock: time.Now,
	}
}

func (s *AlertService) EvaluateRule(ctx context.Context, rule domain.Rule, metricValue float64) error {

	alert, err := s.repo.GetAlertByRule(ctx, string(rule.ID))
	if err != nil {
		return err
	}

	now := s.clock()

	if rule.IsViolated(metricValue) {
		if alert == nil {
			newAlert := domain.NewAlert(domain.AlertID(uuid.New().String()), string(rule.ID), domain.AlertStateActive, rule.Severity, now)
			// TODO: Add logger
			return s.repo.SaveAlert(ctx, newAlert)
		}
		return nil
	}

	if alert != nil {
		alert.Resolve(now)
		return s.repo.SaveAlert(ctx, alert)
	}

	return nil
}
