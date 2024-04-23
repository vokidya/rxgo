package observable

type Subscriber[T any] struct {
	isCompleted bool
	observer    Observer[T]
}

func (s *Subscriber[T]) Next(data T) {
	if !s.isCompleted {
		s.observer.next(data)
	}
}

func (s *Subscriber[T]) Complete() {
	s.isCompleted = true
}

func (s *Subscriber[T]) Error(err any) {
	s.observer.error(err)
}

func NewSubscriber[T any](observer Observer[T]) Subscriber[T] {
	s := Subscriber[T]{}
	s.observer = observer

	return s
}
