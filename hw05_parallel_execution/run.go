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
	schedule := make(chan Task, n)
	quitCh := make(chan struct{})
	errs := &atomic.Int32{}
	errs.Store(int32(m))

	for i := 0; i < n; i++ {
		wg.Add(1)
		go work(wg, schedule, errs, quitCh)
	}

	defer wg.Wait()
	defer close(schedule)
	for _, task := range tasks {
		select {
		case <-quitCh:
			return ErrErrorsLimitExceeded
		case schedule <- task:
		}
	}

	return nil
}

func work(wg *sync.WaitGroup, schedule chan Task, m *atomic.Int32, quitCh chan struct{}) {
	defer wg.Done()
	for task := range schedule {
		err := task()
		if err != nil {
			m.Add(-1)
		}
		if m.Load() <= 0 {
			select {
			case quitCh <- struct{}{}:
			default:
			}
			return
		}
	}
}
