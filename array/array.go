package array

import (
	"encoding/binary"
	"fmt"
)

func array_test() {
	//var variable_name [SIZE] variable_type
	var test_arr [1]int
	test_arr[0] = 1

	fmt.Printf("array:%d", test_arr[0])

	var test_arr1 = [5]float32{9.094, 0.156, 0.349, 8.093, 10.3}

	var j int
	for j = 0; j < len(test_arr1); j++ {
		fmt.Printf("Element[%d]=%f\n", j, test_arr1[j])
	}
	for i, j := range test_arr1 {
		println(i, j)
	}
	fmt.Println(test_arr1[1:])
}

//Go 语言向函数传递数组

func getAverage(arr [5]int, size int) float32 {
	var i, sum int
	var avg float32
	for i = 0; i < size; i++ {
		sum += arr[i]
	}
	avg = float32(sum) / float32(size)
	return avg

}
func slice() {
	var test = [3]int{1, 2, 3}
	fmt.Println(test[1:3])
}
func slicePtr(a []int64) {
	//a=append(a, 2)
	fmt.Println(a)
	b := []int64{2, 4, 5}
	a = append(a, b...)
	fmt.Println(a)
}

func SliceAppendStr() {
	demo := append([]byte("hello "), "world"...)
	fmt.Println(string(demo))
}
func intToByteSlice() {
	c := []byte{67, 78, 90, 89, 93}
	var a uint16 = 17
	binary.BigEndian.PutUint16(c[2:4], a)
	fmt.Print(c)
}
