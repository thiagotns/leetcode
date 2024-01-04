package main

import "fmt"

func removeElement(nums []int, val int) int {

	k := 0
	i := len(nums) - 1
	for i >= k {

		if nums[i] == val {
			i--
			continue
		}

		aux := nums[k]
		nums[k] = nums[i]
		nums[i] = aux

		k++
	}

	fmt.Println(nums)
	return k
}

func main() {

	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2

	r := removeElement(nums, val)

	fmt.Println(r)
}
