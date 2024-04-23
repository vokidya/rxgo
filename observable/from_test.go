package observable

import (
	"fmt"
	"testing"
)

func TestOmit(t *testing.T) {
	t.Parallel()

	dataList := []int{0, 1}

	From(dataList).Subscribe(func(data int) {
		fmt.Println("data", data)
	})
}
