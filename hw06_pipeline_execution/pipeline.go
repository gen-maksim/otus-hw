package hw06pipelineexecution

import (
	"sync"
)

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for _, stage := range stages {
		numChans := 5
		finders := make([]<-chan interface{}, numChans)
		for i := 0; i < numChans; i++ {
			finders[i] = makeDoneStage(in, done, stage)
		}
		in = fanIn(done, finders...)
	}
	return in
}

func fanIn(done In, channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	multiplexedStream := make(chan interface{})
	multiplex := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}

	// Select from all the channels.
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	// Wait for all the reads to complete.
	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}

func makeDoneStage(in In, done In, stage Stage) Out {
	take := make(Bi)
	go func() {
		defer close(take)
		select {
		case <-done:
			return
		default:
			for v := range stage(in) {
				select {
				case <-done:
					return
				default:
					take <- v
				}
			}
		}
	}()
	return take
}
