package operators

import (
	"fmt"
	"rxgo/observable"
	"testing"
)

func TestFrom(t *testing.T) {
	t.Parallel()

	dataList := []int{0, 1, 2, 3, 4}

	observable.Pipe2(observable.From(dataList),
		Map(func(data int, _ int) int {
			return data + 1
		}),
		Map(func(data int, _ int) int {
			return data + 1
		}),
	).Subscribe(func(data int) {
		fmt.Println("data", data)
	}, nil, nil)
}
