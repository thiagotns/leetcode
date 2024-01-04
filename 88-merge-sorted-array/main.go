package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) []int {

	if n == 0 {
		return nums1
	}

	if m == 0 {
		//nums1 = nums2
		copy(nums1, nums2)
		return nums1
	}

	pos := 0

	for i := 0; i < m+n && pos < n; i++ {

		if nums2[pos] >= nums1[i] && i < m+pos {
			continue
		}

		for j := m + n - 1; j > i; j-- {
			nums1[j] = nums1[j-1]
		}

		nums1[i] = nums2[pos]

		pos++
	}

	return nums1
}

func main() {

	nums1 := []int{4, 0, 0, 0, 0, 0}
	m := 1
	nums2 := []int{1, 2, 3, 5, 6}
	n := 5

	r := merge(nums1, m, nums2, n)

	fmt.Println(r)
}
