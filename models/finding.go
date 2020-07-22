package models

import "time"

type Finding struct {
	UUID           string    `json:"uuid"`
	RegressionUUID string    `json:"regression_uuid"`
	DomainToken    string    `json:"domain_token"`
	Signature      string    `json:"signature"`
	URL            string    `json:"url"`
	FoundAt        string    `json:"found_at"`
	StartTimestamp time.Time `json:"start_timestamp"`
	EndTimestamp   time.Time `json:"end_timestamp"`
	Title          string    `json:"title"`
	Definition     struct {
		UUID        string `json:"uuid"`
		Description string `json:"description"`
		Risk        string `json:"risk"`
		References  []struct {
			UUID   string `json:"uuid"`
			Link   string `json:"link"`
			Name   string `json:"name"`
			Source string `json:"source"`
		} `json:"references"`
	} `json:"definition"`
	Score []struct {
		Version string  `json:"version"`
		Score   float64 `json:"score"`
		Vector  string  `json:"vector"`
	} `json:"score"`
	Owasp []struct {
		Year           string `json:"year"`
		Classification string `json:"classification"`
	} `json:"owasp"`
	Details []struct {
		UUID  string `json:"uuid"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"details"`
	Tags []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"tags"`
	Target struct {
		UUID    string `json:"uuid"`
		Type    string `json:"type"`
		Address string `json:"address"`
	} `json:"target"`
}
