// https://leetcode.com/problems/guess-number-higher-or-lower

package guessnumber

var N = 10
var pick = 6

func Input() int {
	return 0
}

func guess(num int) int {
	if num > pick {
		return 1
	} else if num < pick {
		return -1
	}
	return 0
}

func Run(num int) int {
	from := 0
	to := num
	for {
		g := (from + to) / 2
		r := guess(g)
		if r == 0 {
			return g
		}
		if r < 0 {
			to = g - 1
		} else {
			from = g + 1
		}
	}
}
