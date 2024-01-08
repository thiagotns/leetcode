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
			profit += maxProfit(prices[i:])
		}

	}
	return profit
}

func main() {

	prices := []int{1, 2, 3, 4, 5}

	r := maxProfit(prices)

	fmt.Println(r)
}
