// https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum/

package minvaluetogetpositive

func Input() (int, int) {
	return 1, 4
}

func hammingDistance(x int, y int) int {
	ret := 0
	x_xor_y := x ^ y
	for x_xor_y != 0 {
		// ret += x_xor_y & 1
		// x_xor_y >>= 1
		// Brian Kernighan's Algorithm: count 1s bit
		ret++
		x_xor_y &= x_xor_y - 1
	}

	for ; x_xor_y != 0; x_xor_y >>= 1 {
		ret += x_xor_y & 1
	}
	return ret
}

func Run(x, y int) int {
	return hammingDistance(x, y)
}
