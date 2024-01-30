package main

import "fmt"

func twoSum(numbers []int, target int) []int {

	memo := make(map[int]int)

	for i, v := range numbers {

		if val, ok := memo[v]; ok {
			return []int{val, i}
		}

		memo[target-v] = i

	}

	return []int{0, 0}
}

func main() {

	numbers := []int{2, 7, 11, 15}

	r := twoSum(numbers, 9)

	fmt.Println(r)
}
