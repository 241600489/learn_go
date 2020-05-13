package array

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {

	array_test()

}

func TestGetAvg(t *testing.T) {
	var arr = [...]int{8, 8, 8, 8, 0}

	fmt.Println(getAverage(arr, 5))

}
