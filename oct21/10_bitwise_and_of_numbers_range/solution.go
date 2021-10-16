// https://leetcode.com/problems/bitwise-and-of-numbers-range/

package bitwiseandofnumbersrange

func Input() (int, int) {
	return 5, 5
}

func Run(left, right int) int {
	ret := 0
	mask := 1 << 32
	for ; mask > 0; mask >>= 1 {
		if mask&right > 0 {
			break
		}
	}
	x := left ^ right
	for ; mask > 0; mask >>= 1 {
		if mask&x == 0 {
			ret |= mask & left & right
		} else {
			break
		}

	}
	return ret
}
