package observable

type Observable[T any] struct {
	onSubscribe func(subscriber Subscriber[T]) func()
}

func NewObservable[T any](onSubscribe func(subscriber Subscriber[T]) func()) Observable[T] {
	observable := Observable[T]{}
	observable.onSubscribe = onSubscribe
	return observable
}

func (o Observable[T]) Subscribe(next func(data T), complete func(), err func(err any)) Subscription[T] {
	observer := NewObserver[T](next, nil, nil)
	subscriber := NewSubscriber[T](observer)

	onUnsubscribe := o.onSubscribe(subscriber)
	return NewSubscription[T](onUnsubscribe)
}
