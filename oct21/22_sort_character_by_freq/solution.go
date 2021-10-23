// https://leetcode.com/problems/sort-characters-by-frequency/

package sortcharacterbyfreq

import "sort"

func Input() string {
	return "Aabb"
}

func Run(s string) string {
	freq := make(map[byte]int, 52)
	chars := make([]byte, 0, 52)
	for i := range s {
		_, ok := freq[s[i]]
		if !ok {
			freq[s[i]] = 1
			chars = append(chars, s[i])
			continue
		}
		freq[s[i]]++
	}

	sort.Slice(chars, func(i, j int) bool {
		return freq[chars[i]] > freq[chars[j]]
	})

	res := make([]byte, len(s))
	for i := range chars {
		f := freq[chars[i]]
		for j := 0; j < f; j++ {
			res = append(res, chars[i])
		}
	}
	return string(res)
}
