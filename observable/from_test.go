package observable

import (
	"fmt"
	"testing"
)

func TestFrom(t *testing.T) {
	t.Parallel()

	dataList := []int{0, 1, 2, 3, 4}

	From(dataList).Subscribe(func(data int) {
		fmt.Println("data", data)
	}, nil, nil)
}
