package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for _, stage := range stages {
		in = makeDoneStage(in, done, stage)
	}

	return in
}

func makeDoneStage(in In, done In, stage Stage) Out {
	take := make(Bi)
	go func() {
		defer close(take)
		out := stage(in)
		for {
			select {
			case <-done:
				return
			case v, ok := <-out:
				if !ok {
					return
				}
				take <- v
			}
		}
	}()
	return take
}
