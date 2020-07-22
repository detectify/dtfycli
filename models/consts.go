package models

type Severity int

const (
	SeverityAny Severity = iota
	SeverityHigh
	SeverityMedium
	SeverityLow
	SeverityInformation
)
