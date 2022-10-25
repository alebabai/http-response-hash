package pool

import (
	"fmt"
)

type Pool[T, V any] struct {
	action   Action[T, V]
	consumer Consumer[V]
	size     int
}

type Action[T, V any] func(in T) V

type Consumer[T any] func(in T)

func New[T, V any](
	a Action[T, V],
	c Consumer[V],
	size int,
) (*Pool[T, V], error) {
	p := &Pool[T, V]{
		action:   a,
		consumer: c,
		size:     size,
	}

	if err := p.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate a pool: %w", err)
	}

	return p, nil
}

func (p *Pool[T, V]) Process(values ...T) {
	in := make(chan T, len(values))
	out := make(chan V, len(values))

	for i := 0; i < p.size; i++ {
		go work(in, out, p.action)
	}

	for _, v := range values {
		in <- v
	}
	close(in)

	for i := 0; i < len(values); i++ {
		p.consumer(<-out)
	}
	close(out)
}

func work[T, V any](in <-chan T, out chan<- V, a Action[T, V]) {
	for v := range in {
		out <- a(v)
	}
}
