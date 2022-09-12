package pool

type Action[T, V any] func(in T) V

func Worker[T, V any](in <-chan T, out chan<- V, a Action[T, V]) {
	for v := range in {
		out <- a(v)
	}
}

type Consumer[T any] func(in T)

type Pool[T, V any] struct {
	Action   Action[T, V]
	Consumer Consumer[V]
	Size     int
}

func (p *Pool[T, V]) Process(values ...T) {
	in := make(chan T, len(values))
	out := make(chan V, len(values))

	for i := 0; i < p.Size; i++ {
		go Worker(in, out, p.Action)
	}

	for _, v := range values {
		in <- v
	}
	close(in)

	for i := 0; i < len(values); i++ {
		p.Consumer(<-out)
	}
	close(out)
}
