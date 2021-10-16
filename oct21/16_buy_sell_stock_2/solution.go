// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/

package buysellstock2

func Input() []int {
	// return []int{3, 3, 5, 0, 0, 3, 1, 4}
	// return []int{1, 2, 3, 4, 5}
	// return []int{7, 6, 4, 3, 1}
	// return []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}
	// return []int{6, 1, 6, 4, 3, 0, 2}
	return []int{1, 2, 3, 4, 5}
}

func Min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func Max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func Run(prices []int) int {
	profit := 0
	l := len(prices)

	// calculate profit of first part
	profit_1 := make([]int, l)
	profit_1[0] = 0
	min := prices[0]
	for i := 1; i < l-1; i++ {
		profit_1[i] = Max(profit_1[i-1], prices[i]-min)
		min = Min(min, prices[i])
	}

	max, profit_2 := prices[l-1], 0
	for i := l - 2; i >= 0; i-- {
		profit_2 = Max(profit_2, max-prices[i])
		max = Max(max, prices[i])
		if profit_1[i]+profit_2 > profit {
			profit = profit_1[i] + profit_2
		}
	}
	return profit
}
