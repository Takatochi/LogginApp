package models

import "time"

type LogEntry struct {
	Level     string    `json:"Level"`
	Category  string    `json:"Category"`
	Message   string    `json:"Message"`
	Exception string    `json:"Exception,omitempty"`
	Timestamp time.Time `json:"Timestamp"`
}
