// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/

package buysellstock

func Input() []int {
	return []int{1, 2, 3, 0, 2}
}

func MaxInt(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func MaxArr(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}

//
// func Run(prices []int) int {
//   track := make([][]int, len(prices))
//   for i := range track {
//     track[i] = []int{0, 0, 0}
//   }
//   // use track 3 scenerios
//   // 0 : buy
//   // 1 : sell
//   // 2 : in cool down
//   track[0][0] = -prices[0]
//   for i := 1; i < len(prices); i++ {
//     track[i][0] = MaxInt(track[i-1][0], track[i-1][2]-prices[i]) // get last balance cooldown
//     track[i][1] = MaxInt(track[i-1][1], track[i-1][0]+prices[i]) // get the last buy
//     track[i][2] = MaxInt(track[i-1][2], track[i-1][1])
//   }
//   return MaxArr(track[len(prices)-1])
// }

type Scenerio struct {
	buy, sell, cooldown int
}

func Run(prices []int) int {
	track := Scenerio{
		buy:      -prices[0],
		sell:     0,
		cooldown: 0,
	}
	for i := 1; i < len(prices); i++ {
		last_buy, last_sell := track.buy, track.sell
		track.buy = MaxInt(last_buy, track.cooldown-prices[i])
		track.sell = MaxInt(last_sell, last_buy+prices[i])
		track.cooldown = MaxInt(track.cooldown, last_sell)
	}
	res := track.buy
	if track.sell > res {
		res = track.sell
	}
	if track.cooldown > res {
		res = track.cooldown
	}
	return res
}
