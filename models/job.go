package models

import "time"

type Job struct {
	ID       string
	Interval time.Duration
	Task     func()
}
