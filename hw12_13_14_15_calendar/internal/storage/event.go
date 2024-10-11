package storage

import "time"

type Event struct {
	ID        string
	Title     string
	StartTime time.Time
	EndTime   time.Time
	UserId    int
}
