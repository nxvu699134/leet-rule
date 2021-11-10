// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/

package buysellstock2

func Input() []int {
	return []int{1, 2, 3, 0, 2}
}

func MaxInt(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

type Scenerio struct {
	buy, sell int
}

func maxProfit(prices []int) int {
	s := Scenerio{
		buy:  -prices[0],
		sell: 0,
	}
	for i := 1; i < len(prices); i++ {
		last_buy, last_sell := s.buy, s.sell
		s.buy = MaxInt(last_buy, last_sell-prices[i])
		s.sell = MaxInt(last_sell, last_buy+prices[i])
	}
	return MaxInt(s.buy, s.sell)
}

func Run(prices []int) int {
	return maxProfit(prices)
}
