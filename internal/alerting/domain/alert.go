package domain

import "time"

type AlertID string

type Alert struct {
	ID         AlertID       `bson:"id,omitempty"`
	RuleID     string        `bson:"rule_id,omitempty"`
	State      AlertState    `bson:"state,omitempty"`
	Severity   AlarmSeverity `bson:"severity,omitempty"`
	CreatedAt  time.Time     `bson:"created_at"`
	ResolvedAt *time.Time    `bson:"resolved_at,omitempty"`
	IgnoredAt  *time.Time    `bson:"ignored_at,omitempty"`
}

func NewAlert(id AlertID, ruleId string, state AlertState, severity AlarmSeverity, time time.Time) *Alert {
	return &Alert{
		ID:        id,
		RuleID:    ruleId,
		State:     state,
		Severity:  severity,
		CreatedAt: time,
	}
}

func (a *Alert) Resolve(time time.Time) {
	if a.State == AlertStateResolved {
		return
	}

	a.State = AlertStateResolved
	a.ResolvedAt = &time
}

func (a *Alert) Ignore(time time.Time) {
	if a.State != AlertStateActive {
		return
	}

	a.State = AlertStateIgnored
	a.IgnoredAt = &time
}
