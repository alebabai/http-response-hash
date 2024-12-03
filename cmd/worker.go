package cmd

import (
	"sync"
)

type WorkerPool[T, V any] struct {
	Action   func(in T) V
	Consumer func(in V)
	Size     uint
}

func (p *WorkerPool[T, V]) Process(values ...T) {
	if p.Size == 0 {
		return
	}

	var wg sync.WaitGroup

	in := make(chan T, len(values))
	out := make(chan V, len(values))

	for i := 0; i < int(p.Size); i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for v := range in {
				out <- p.Action(v)
			}
		}()
	}

	go func() {
		defer close(in)

		for _, v := range values {
			in <- v
		}
	}()

	go func() {
		defer close(out)

		wg.Wait()
	}()

	for result := range out {
		p.Consumer(result)
	}
}
