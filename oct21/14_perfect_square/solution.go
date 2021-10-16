// https://leetcode.com/problems/perfect-squares/

package perfectsquare

func Input() int {
	return 12
}

func Bfs(input, queue *[]int, num int, depth int) int {
	new_queue := make([]int, 0)
	for i := range *queue {
		accum := (*queue)[i]
		for j := range *input {
			new_accum := accum + (*input)[j]
			if new_accum > num {
				continue
			} else if new_accum == num {
				return depth + 1
			} else {
				new_queue = append(new_queue, new_accum)
			}
		}
	}
	if len(new_queue) == 0 {
		return -1
	}
	return Bfs(input, &new_queue, num, depth+1)
}

func Run(num int) int {
	if num == 1 {
		return 1
	}
	sq := make([]int, 0)
	for i := 1; ; i++ {
		s := i * i
		if s == num {
			return 1
		}
		if s > num {
			break
		}
		sq = append([]int{s}, sq...)
	}

	queue := make([]int, len(sq))
	copy(queue, sq)

	return Bfs(&sq, &queue, num, 0) + 1
}
