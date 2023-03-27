package schedulers

import "time"

// Scheduler
type Scheduler struct {
}

// NewScheduler
func NewScheduler() *Scheduler {
	return &Scheduler{}
}

func Attach(job func(), interval time.Duration) (string, error) {
	// TODO attach job to scheduler
	return "", nil
}

func Detach(id string) {
	// TODO detach job from scheduler
}
