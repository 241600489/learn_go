package builtn

import (
	"fmt"
	"log"
)

func copySlice() {
	nums := []int{1, 2, 3, 4}
	dest := make([]int, 3)
	dest[0] = 1
	copyCount := copy(dest[1:], nums[2:3])
	log.Printf("copy count:%d", copyCount)
	nums[3] = 0
	fmt.Println(nums, dest)
	fmt.Println(nums[2:])

}
