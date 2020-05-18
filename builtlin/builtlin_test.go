package builtlin

import (
	"fmt"
	"testing"
)

func TestDeferFirst(t *testing.T) {
	result := f1()
	fmt.Println("result's result is ", result)

	result2 := f2()

	fmt.Println("f2's result is ", result2)
}
