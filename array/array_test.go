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

func TestSlice(t *testing.T) {
	slice()
}

func TestSlicePtr(t *testing.T) {
	a := []int64{1, 3}
	slicePtr(a)
}
func TestSliceAppendStr(t *testing.T) {
	SliceAppendStr()
}

func TestIntToByteSlice(t *testing.T) {
	intToByteSlice()

}
