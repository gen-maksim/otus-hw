package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	in = makeDoneStage(done, in)
	for _, stage := range stages {
		in = makeDoneStage(done, stage(in))
	}

	return in
}

func makeDoneStage(done In, out Out) Out {
	take := make(Bi)
	go func() {
		defer close(take)
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
