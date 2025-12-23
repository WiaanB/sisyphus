package domain

type AlarmSeverity string

const (
	AlarmSeverityCritical AlarmSeverity = "critical"
	AlarmSeverityHigh     AlarmSeverity = "high"
	AlarmSeverityMedium   AlarmSeverity = "medium"
	AlarmSeverityLow      AlarmSeverity = "low"
)
