package domain

type RuleID string

type Rule struct {
	ID        RuleID
	Metric    string
	Threshold float64
	Severity  AlarmSeverity
}

func (r Rule) IsViolated(value float64) bool {
	return value > r.Threshold
}
