package events

import "time"

type Event struct {
	Id        int       `json:"id"`
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
	Odometer  int       `json:"odometer"`
}
