package observable

import "fmt"

func From[T any](data []T) Observable[T] {
	ch := make(chan T)

	onSubscribe := func(subscriber Subscriber[T]) func() {
		go func() {
			for item := range ch {
				subscriber.Next(item)
			}
		}()

		for _, item := range data {
			ch <- item
		}

		return func() {
			fmt.Println("This is unsbscriped")
		}
	}

	observable := NewObservable[T](onSubscribe)

	return observable
}
