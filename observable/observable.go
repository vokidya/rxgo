package observable

type Observable[T any] struct {
	onSubscribe func(subscriber Subscriber[T]) func()
}

func (o Observable[T]) Subscribe(next func(data T)) Subscription[T] {
	observer := NewObserver[T](next, func() {})
	subscriber := NewSubscriber[T](observer)

	onUnsubscribe := o.onSubscribe(subscriber)
	return NewSubscription[T](onUnsubscribe)
}

func NewObservable[T any](onSubscribe func(subscriber Subscriber[T]) func()) Observable[T] {
	observable := Observable[T]{}
	observable.onSubscribe = onSubscribe
	return observable
}
