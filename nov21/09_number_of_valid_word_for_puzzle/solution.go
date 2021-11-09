// https://leetcode.com/problems/unique-binary-search-trees/

package uniquebinsearchtrees

func Input() ([]string, []string) {
	return []string{"aaaa", "asas", "able", "ability", "actt", "actor", "access"},
		[]string{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"}
}

//
// //this solution is not satified time limit
// func findNumOfValidWords(words []string, puzzles []string) []int {
//   ret := make([]int, len(puzzles))
//   for i := 0; i < len(words); i++ {
//     m := make(map[byte]struct{})
//     for j := 0; j < len(words[i]); j++ {
//       m[words[i][j]] = struct{}{}
//     }
//     if len(m) > 7 { // this is because all len of puzzle are 7
//       continue
//     }
//     for j := 0; j < len(puzzles); j++ {
//       if _, is_contain_first_letter := m[puzzles[j][0]]; !is_contain_first_letter {
//         continue
//       }
//       count := 1
//       for k := 1; k < 7; k++ {
//         if _, contain := m[puzzles[j][k]]; contain {
//           count++
//         }
//       }
//       if count == len(m) {
//         ret[j]++
//       }
//     }
//   }
//   return ret
// }

//========================================================================================
// This solution satified time limit but still consider very slow
// func IntersectArray(arr1, arr2 []int) []int {
//   i, j := 0, 0
//   ret := make([]int, 0)
//   for i < len(arr1) && j < len(arr2) {
//     if arr1[i] < arr2[j] {
//       i++
//     } else if arr1[i] > arr2[j] {
//       j++
//     } else {
//       if len(ret) > 0 && arr1[i] == ret[len(ret)-1] {
//         i++
//         j++
//       } else {
//         ret = append(ret, arr1[i])
//         i++
//         j++
//       }
//     }
//   }
//   return ret
// }
//
// func findNumOfValidWords(words []string, puzzles []string) []int {
//   ret := make([]int, len(puzzles))
//
//   contain_map := make(map[byte][]int)
//   for i := 0; i < len(puzzles); i++ {
//     for j := 0; j < len(puzzles[i]); j++ {
//       if _, ok := contain_map[puzzles[i][j]]; !ok {
//         contain_map[puzzles[i][j]] = make([]int, 0)
//       }
//       contain_map[puzzles[i][j]] = append(contain_map[puzzles[i][j]], i)
//     }
//   }
//
//   for i := 0; i < len(words); i++ {
//     m := make(map[byte]struct{})
//     letters := make([]byte, 0)
//     for j := 0; j < len(words[i]); j++ {
//       if _, ok := m[words[i][j]]; !ok {
//         m[words[i][j]] = struct{}{}
//         letters = append(letters, words[i][j])
//       }
//     }
//     if len(m) > 7 { // this is because all len of puzzle are 7
//       continue
//     }
//
//     idx_arr := contain_map[letters[0]]
//     for j := 1; j < len(letters); j++ {
//       idx_arr = IntersectArray(idx_arr, contain_map[letters[j]])
//     }
//     for j := 0; j < len(idx_arr); j++ {
//       if _, ok := m[puzzles[idx_arr[j]][0]]; ok {
//         ret[idx_arr[j]]++
//       }
//     }
//   }
//   return ret
// }
//=====================================================================================

func makeBitsFromString(s *string) int {
	ret := 0
	for i := 0; i < len(*s); i++ {
		ret |= 1 << ((*s)[i] - 'a')
	}
	return ret
}

func findNumOfValidWords(words []string, puzzles []string) []int {
	ret := make([]int, len(puzzles))
	m := make(map[int]int)
	for i := 0; i < len(words); i++ {
		b := makeBitsFromString(&words[i])
		if _, ok := m[b]; ok {
			m[b]++
		} else {
			m[b] = 1
		}
	}

	// This version is "BIT" version of the first solution => better memory used
	// in edge case like len(words) = 10^5
	// with each puzzle we loop 10^5 time
	// but because of every single puzzle is only 7 and letters in puzzle are unique
	// so we can generate 2^7 =128 combination from each puzzle
	// with every single combination is a sub set of puzzle
	// ref: https://cp-algorithms.com/algebra/all-submasks.html

	for i := 0; i < len(puzzles); i++ {
		p := makeBitsFromString(&puzzles[i])
		first := 1 << (puzzles[i][0] - 'a')
		for submask := p; submask != 0; submask = (submask - 1) & p {
			if submask&first != 0 {
				ret[i] += m[submask]
			}
		}
	}

	for i := 0; i < len(puzzles); i++ {
		p := makeBitsFromString(&puzzles[i])
		first := 1 << (puzzles[i][0] - 'a')
		for b, freq := range m {
			if b&first == 0 {
				continue
			}
			if b & ^p == 0 {
				ret[i] += freq
			}
		}
	}
	return ret
}

func Run(words []string, puzzles []string) []int {
	// a := makeBitsFromString(&puzzles[0])
	// fmt.Printf("%032b\n", a)
	// for submask := a; submask != 0; submask = (submask - 1) & a {
	//   fmt.Printf("%032b\n", submask)
	// }
	return findNumOfValidWords(words, puzzles)
}
