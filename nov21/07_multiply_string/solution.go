// https://leetcode.com/problems/multiply-strings/

package multiplystring

func Input() (string, string) {
	return "123", "456"
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	ret := make([]byte, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		n1 := num1[i] - '0'
		for j := len(num2) - 1; j >= 0; j-- {
			ret[i+j+1] += n1 * (num2[j] - '0')
			// carry to next digit
			ret[i+j] += ret[i+j+1] / 10
			ret[i+j+1] %= 10
		}
	}
	start := 0
	for ; start < len(ret); start++ {
		if ret[start] != 0 {
			break
		}
	}
	for i := start; i < len(ret); i++ {
		ret[i] += '0'
	}
	return string(ret[start:])
}

func Run(num1, num2 string) string {
	return multiply(num1, num2)
}
