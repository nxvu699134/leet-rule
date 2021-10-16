// https://leetcode.com/problems/longest-common-subsequence/
// Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.
//
// A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.
//
// For example, "ace" is a subsequence of "abcde".
// A common subsequence of two strings is a subsequence that is common to both strings.

package longestcommonsubsequence

func Input() (string, string) {
	return "abcde", "ace"
}

func Run(txt1, txt2 string) int {
	const BUF_SIZE = 3
	buf := [BUF_SIZE]int{0, 1, 1}
	buf_start := 0
	if n < BUF_SIZE {
		return buf[n]
	}
	for i := BUF_SIZE; i <= n; i++ {
		buf[buf_start] = buf[0] + buf[1] + buf[2]
		buf_start = (buf_start + 1) % BUF_SIZE
	}
	return buf[(buf_start+BUF_SIZE-1)%BUF_SIZE]
}
