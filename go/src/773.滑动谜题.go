/*
 * @lc app=leetcode.cn id=773 lang=golang
 *
 * [773] 滑动谜题
 */

// @lc code=start
func slidingPuzzle(board [][]int) int {
	step := 0
	cur := make([]byte, 6)
	target := "123450"
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			cur[i*3+j] = byte(board[i][j]) + '0'
		}
	}

	// 一维byte切片中，每个位置的元素的邻居在二维board位置 转换到一维byte后的下标
	neighbor := [][]int{{1, 3}, {0, 4, 2}, {1, 5}, {0, 4}, {1, 3, 5}, {2, 4}}
	queue := []string{string(cur)}
	m := map[string]struct{}{}
	for len(queue) > 0 {
		newQueue := []string{}
		for _, str := range queue {
			if str == target {
				return step
			}

			idx := getZeroIndex(str)
			for _, another := range neighbor[idx] {
				sb := []byte(str)
				sb[idx], sb[another] = sb[another], sb[idx]

				nstr := string(sb)
				if _, ok := m[nstr]; ok {
					continue
				}
				m[nstr] = struct{}{}
				newQueue = append(newQueue, nstr)
			}
		}

		queue = newQueue
		step++
	}

	return -1
}

func getZeroIndex(str string) int {
	for i := 0; i < len(str); i++ {
		if str[i] == '0' {
			return i
		}
	}

	return 0
}

// @lc code=end

