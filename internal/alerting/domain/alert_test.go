package domain

import (
	"testing"
	"time"
)

func TestAlert_Resolve(t *testing.T) {
	tests := []struct {
		name     string
		alert    Alert
		resolved bool
	}{
		{
			name: "resolves the alert",
			alert: Alert{
				ID:         AlertID("1"),
				RuleID:     "rule1",
				State:      AlertStateActive,
				Severity:   AlarmSeverityLow,
				CreatedAt:  time.Time{},
				ResolvedAt: nil,
			},
			resolved: true,
		},
		{
			name: "does not resolve the alert",
			alert: Alert{
				ID:         AlertID("1"),
				RuleID:     "rule1",
				State:      AlertStateResolved,
				Severity:   AlarmSeverityLow,
				CreatedAt:  time.Time{},
				ResolvedAt: nil,
			},
			resolved: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.alert.Resolve(time.Time{})

			if test.resolved && test.alert.ResolvedAt == nil {
				t.Error("alert resolved at should not be nil")
			}
		})
	}
}

func TestAlert_Ignore(t *testing.T) {
	tests := []struct {
		name    string
		alert   Alert
		ignored bool
	}{
		{
			name: "ignores the alert",
			alert: Alert{
				ID:        AlertID("1"),
				RuleID:    "rule1",
				State:     AlertStateActive,
				Severity:  AlarmSeverityLow,
				CreatedAt: time.Time{},
				IgnoredAt: nil,
			},
			ignored: true,
		},
		{
			name: "does not ignore the alert",
			alert: Alert{
				ID:        AlertID("1"),
				RuleID:    "rule1",
				State:     AlertStateResolved,
				Severity:  AlarmSeverityLow,
				CreatedAt: time.Time{},
				IgnoredAt: nil,
			},
			ignored: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.alert.Ignore(time.Time{})

			if test.ignored && test.alert.IgnoredAt == nil {
				t.Error("alert resolved at should not be nil")
			}
		})
	}
}
