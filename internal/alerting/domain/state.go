package domain

type AlertState = string

const (
	AlertStateActive   AlertState = "active"
	AlertStateResolved AlertState = "resolved"
	AlertStateIgnored  AlertState = "ignored"
)
