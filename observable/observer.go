package observable

type Observer[T any] struct {
	next     func(data T)
	complete func()
	error    func(err any)
}

func NewObserver[T any](next func(data T), complete func(), error func(err any)) Observer[T] {
	observer := Observer[T]{}

	if next == nil {
		observer.next = func(data T) {}
	} else {
		observer.next = next
	}

	if complete == nil {
		observer.complete = func() {}
	} else {
		observer.complete = complete
	}

	if error == nil {
		observer.error = func(err any) {}
	} else {
		observer.error = error
	}

	return observer
}
