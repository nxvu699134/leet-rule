// https://leetcode.com/problems/longest-duplicate-substring/

package longestdupsubstring

func Input() string {
	return "nyvzwttxsshphczjjklqniaztccdrawueylaelkqtjtxdvutsewhghcmoxlvqjktgawwgpytuvoepnyfbdywpmmfukoslqvdrkuokxcexwugogcwvsuhcziwuwzfktjlhbiuhkxcreqrdbj"
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

func checkExistDupSubString(si *string, buffer_size int, result *string) bool {
	s := *si
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
			for j := range arr {
				k := 0
				for ; k < buffer_size; k++ {
					if s[i+k] != s[arr[j]+k] {
						break
					}
				}
				if k == buffer_size {
					*result = s[i : i+buffer_size]
					return true
				}
			}
			m[new_h] = append(m[new_h], i)
		} else {
			m[new_h] = []int{i}
		}
		h = new_h
	}
	return false
}

func Run(s string) string {
	// the worst case is with long string and max dup sub is small
	res := ""
	l, r := 0, len(s)-1
	for l <= r {
		mid := (l + r) / 2
		if checkExistDupSubString(&s, mid+1, &res) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return res
}

// func checkExistDupSubString(si *string, buffer_size int, result *string) bool {
//   s := *si
//   m := make(map[string]struct{})
//   n := len(s) - buffer_size
//   for i := 0; i <= n; i++ {
//     sub := s[i : i+buffer_size]
//     _, ok := m[sub]
//     if ok {
//       *result = sub
//       return true
//     }
//     m[sub] = struct{}{}
//   }
//   return false
// }
// func Run(s string) string {
//   // naive solution
//   buffer_size := len(s) - 1
//   for ; buffer_size > 0; buffer_size-- {
//   }
//
//   res := ""
//   l, r := 0, len(s)-1
//   for l <= r {
//     mid := (l + r) / 2
//     if checkExistDupSubString(&s, mid+1, &res) {
//       l = mid + 1
//     } else {
//       r = mid - 1
//     }
//   }
//   return res
// }
