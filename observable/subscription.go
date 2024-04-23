package observable

type Subscription[T any] struct {
	onUnSubscribe func()
}

func (s *Subscription[T]) UnSubscribe() {
	s.onUnSubscribe()
}

func NewSubscription[T any](unsubscrine func()) Subscription[T] {
	subscription := Subscription[T]{}
	subscription.onUnSubscribe = unsubscrine

	return subscription
}
