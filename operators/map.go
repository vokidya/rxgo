package operators

import "rxgo/observable"

func Map[T any, R any](project func(_value T, _index int) R) func(data observable.Observable[T]) observable.Observable[R] {
	return func(source observable.Observable[T]) observable.Observable[R] {
		return observable.NewObservable(func(subscriber observable.Subscriber[R]) func() {
			index := 0
			subscription := source.Subscribe(func(sourceValue T) {
				index++
				projectedValue := project(sourceValue, index)

				subscriber.Next(projectedValue)
			}, subscriber.Complete, subscriber.Error)

			return subscription.UnSubscribe
		})
	}
}
