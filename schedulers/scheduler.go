package schedulers

import (
	"fmt"
	"sync"
	"time"

	"github.com/Bulut-Bilisimciler/certman/models"
)

// Scheduler
type Scheduler struct {
	jobs     []*models.Job
	stopChan chan struct{}
	wg       sync.WaitGroup
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		stopChan: make(chan struct{}),
	}
}

func (s *Scheduler) Attach(task func(), interval time.Duration) (string, error) {

	// Create a new job with a unique ID
	id := fmt.Sprintf("%d", time.Now().UnixNano())
	job := &models.Job{
		ID:       id,
		Task:     task,
		Interval: interval,
	}

	s.jobs = append(s.jobs, job)

	return job.ID, nil
}

func (s *Scheduler) Detach(id string) {
	for i, j := range s.jobs {
		if j.ID == id {
			s.jobs = append(s.jobs[:i], s.jobs[i+1:]...)
			break
		}
	}
}

func (s *Scheduler) Start() {
	s.wg.Add(len(s.jobs))

	go func() {
		for _, j := range s.jobs {
			go func(j *models.Job) {
				ticker := time.NewTicker(j.Interval)
				defer ticker.Stop()

				for {
					select {
					case <-s.stopChan:
						s.wg.Done()
						return
					case <-ticker.C:
						j.Task()
					}
				}
			}(j)
		}
	}()
}

func (s *Scheduler) Stop() {
	close(s.stopChan)
	s.wg.Wait()
}
