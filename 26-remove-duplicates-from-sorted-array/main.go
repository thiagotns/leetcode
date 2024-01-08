package main

import "fmt"

func removeDuplicates(nums []int) int {

	c := 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			nums[c] = nums[i]
			c++
		}
	}

	nums[c] = nums[len(nums)-1]
	c++

	fmt.Println(nums)
	return c

}

func main() {

	nums := []int{1, 1, 2}
	r := removeDuplicates(nums)

	fmt.Println(r)

}
