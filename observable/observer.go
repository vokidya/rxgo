package observable

type Observer[T any] struct {
	next     func(data T)
	complete func()
}

func NewObserver[T any](next func(data T), complete func()) Observer[T] {
	observer := Observer[T]{}
	observer.next = next
	observer.complete = complete

	return observer
}
