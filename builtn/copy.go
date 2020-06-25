package builtn

import (
	"fmt"
	"log"
)

func copySlice() {
	nums := []int{1, 2, 3, 4}
	dest := make([]int, len(nums))
	copyCount := copy(dest, nums)
	log.Printf("copy count:%d", copyCount)
	nums[3] = 0
	fmt.Println(nums, dest)

}
