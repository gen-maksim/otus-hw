package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	wg := &sync.WaitGroup{}
	schedule := make(chan Task)
	errs := &atomic.Int32{}
	errs.Store(int32(m))

	for i := 0; i < n; i++ {
		wg.Add(1)
		go work(wg, schedule, errs)
	}

	var err error
	for _, task := range tasks {
		load := errs.Load()
		if load <= 0 {
			err = ErrErrorsLimitExceeded
			break
		}
		schedule <- task
	}

	close(schedule)
	wg.Wait()

	return err
}

func work(wg *sync.WaitGroup, schedule chan Task, m *atomic.Int32) {
	for task := range schedule {
		err := task()
		if err != nil {
			m.Add(-1)
		}
	}

	wg.Done()
}
