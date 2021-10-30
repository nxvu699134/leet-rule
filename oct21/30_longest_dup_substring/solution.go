// https://leetcode.com/problems/longest-duplicate-substring/

package longestdupsubstring

func Input() string {
	return "banana"
}

const (
	BASE  = 32
	PRIME = 101
)

func hash(s string) int {
	res := 0
	for i := range s {
		res = (res*BASE + int(s[i]-'a')) % PRIME
	}
	return res
}

func Run(s string) string {
	buffer_size := len(s) - 1
	for ; buffer_size > 0; buffer_size-- {
		m := make(map[int][]int)
		n := len(s) - buffer_size
		h := hash(s[0:buffer_size])
		m[h] = []int{0}
		w := 1
		for i := 0; i < buffer_size-1; i++ {
			w = (w % PRIME) * BASE
		}
		for i := 1; i <= n; i++ {
			new_h := ((h+PRIME-int(s[i-1]-'a')*w%PRIME)*BASE + int(s[i+buffer_size-1]-'a')) % PRIME
			if arr, ok := m[new_h]; ok {
				sub := s[i : i+buffer_size]
				for j := range arr {
					cmp_sub := s[arr[j] : arr[j]+buffer_size]
					if sub == cmp_sub {
						return sub
					}
				}
				m[new_h] = append(m[new_h], i)
			} else {
				m[new_h] = []int{i}
			}
			h = new_h
		}
	}
	return ""
}

// func Run(s string) string {
//   // naive solution
//   buffer_size := len(s) - 1
//   for ; buffer_size > 0; buffer_size-- {
//     m := make(map[string]struct{})
//     n := len(s) - buffer_size
//     for i := 0; i <= n; i++ {
//       sub := s[i : i+buffer_size]
//       _, ok := m[sub]
//       if ok {
//         return sub
//       }
//       m[sub] = struct{}{}
//     }
//   }
//   return ""
// }
