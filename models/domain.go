package models

import "time"

type Domain struct {
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
	Token     string    `json:"token"`
	Monitored bool      `json:"monitored"`
	Owner     struct {
		Name string `json:"name"`
	} `json:"owner"`
}
