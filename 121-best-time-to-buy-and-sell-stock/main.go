package main

import "fmt"

func maxProfit(prices []int) int {
	profit := 0
	min := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
			continue
		}

		if prices[i]-min > profit {
			profit = prices[i] - min
		}

	}
	return profit
}

func main() {

	prices := []int{7, 1, 5, 3, 6, 4}

	r := maxProfit(prices)

	fmt.Println(r)
}
