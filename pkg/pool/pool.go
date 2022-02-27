package pool

type WorkerAction func(in interface{}) string

func Worker(in <-chan interface{}, out chan<- interface{}, action WorkerAction) {
	for v := range in {
		out <- action(v)
	}
}

type WorkerConsumer func(out interface{})

type WorkerPool struct {
	Action   WorkerAction
	Consumer WorkerConsumer
	Size     int
}

type WorkerPoolOption func(pool *WorkerPool)

func NewWorkerPool(
	action WorkerAction,
	consumer WorkerConsumer,
	opts ...WorkerPoolOption,
) *WorkerPool {
	pool := &WorkerPool{
		Action:   action,
		Consumer: consumer,
	}

	for _, opt := range opts {
		opt(pool)
	}

	return pool
}

func WithSize(size int) WorkerPoolOption {
	return func(pool *WorkerPool) {
		pool.Size = size
	}
}

func (pool *WorkerPool) Process(values ...interface{}) {
	in := make(chan interface{}, len(values))
	out := make(chan interface{}, len(values))

	for i := 0; i < pool.Size; i++ {
		go Worker(in, out, pool.Action)
	}

	for _, v := range values {
		in <- v
	}
	close(in)

	for i := 0; i < len(values); i++ {
		pool.Consumer(<-out)
	}
	close(out)
}
