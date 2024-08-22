package main

import "fmt"

func main() {
	input1 := []int{7, 1, 5, 3, 6, 4}
	input2 := []int{1, 2, 3, 4, 5}
	input3 := []int{7, 6, 4, 3, 1}

	output1 := maxProfit(input1) // 7
	// Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4. Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3
	output2 := maxProfit(input2) // 4
	output3 := maxProfit(input3) // 0

	fmt.Println(output1, output2, output3)
	fmt.Println(output3)

}

func maxProfit(prices []int) int {
	var profit int
	lastPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		price := prices[i]
		if price > lastPrice {
			profit += price - lastPrice
		}
		lastPrice = price
	}

	return profit
}
