package integer

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {

	sum := Sum(4, 4)

	expected := 8

	if sum != expected {
		t.Error("expected ", expected, " sum", sum)
	}
}

func ExampleSum() {
	sum := Sum(1, 5)
	fmt.Println(sum)
	// Output: 6
}
