package hw05parallelexecution

import (
	"errors"
	"sync"
	"time"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type ErrC struct {
	mu sync.Mutex
	i  int
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	wg := &sync.WaitGroup{}
	schedule := make(chan Task, n)
	quitCh := make(chan struct{})
	errC := &ErrC{i: m}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go work(wg, schedule, errC, quitCh)
	}

	defer wg.Wait()
	defer close(schedule)

	taskI := 0
	for {
		select {
		case <-quitCh:
			return ErrErrorsLimitExceeded
		case schedule <- tasks[taskI]:
			taskI++
			if taskI == len(tasks) {
				return nil
			}
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func work(wg *sync.WaitGroup, schedule chan Task, m *ErrC, quitCh chan struct{}) {
	defer wg.Done()
	for {
		task, ok := <-schedule
		if !ok {
			break
		}
		err := task()
		m.mu.Lock()
		if err != nil {
			m.i--
		}
		if m.i <= 0 {
			select {
			case <-quitCh:
			default:
				close(quitCh)
			}
			m.mu.Unlock()
			return
		}
		m.mu.Unlock()
	}
}
