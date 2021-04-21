package runner

import (
	"time"
)

// Job interfact to runner package
type Job interface {
	Run()
}

// JobQueue is slice of Jobs
type JobQueue []Job

// Runner struct to hold JoeQueues
type Runner struct {
	jobQueues map[time.Duration]JobQueue
}

// NewRunner start a new running
func NewRunner() *Runner {
	return &Runner{jobQueues: make(map[time.Duration]JobQueue)}
}

// Schedule create a schedule
func (r *Runner) Schedule(job Job, duration time.Duration) {
	r.jobQueues[duration] = append(r.jobQueues[duration], job)
}

// Run starts the runner
func (r *Runner) Run() {
	for duration, jobQueue := range r.jobQueues {
		ticker := time.NewTicker(duration)

		go func() {
			for ; true; <-ticker.C {
				for _, job := range jobQueue {
					job.Run()
				}
			}
		}()
	}
}
