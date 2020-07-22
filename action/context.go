package action

import "github.com/detectify/dtfycli/models"

type Context struct {
	PrintAsJSON       bool
	GroupByHeuristics bool
	QueryBySeverity   models.Severity
	QueryDaysBack     int
}
