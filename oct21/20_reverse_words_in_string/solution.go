// https://leetcode.com/problems/reverse-words-in-a-string/

package reversewordsinastring

func Input() string {
	return "  hello world  "
}

func Run(s string) string {
	words := make([]string, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		j := i + 1
		for ; j < len(s); j++ {
			if s[j] == ' ' {
				break
			}
		}
		words = append(words, s[i:j])
		i = j
	}
	res := words[len(words)-1]
	for i := len(words) - 2; i >= 0; i-- {
		res += " " + words[i]
	}
	return res
}
