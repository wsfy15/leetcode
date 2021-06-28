/*
 * @lc app=leetcode.cn id=909 lang=golang
 *
 * [909] 蛇梯棋
 */

// @lc code=start
func snakesAndLadders(board [][]int) int {
	// 将二维board转换为一维
	n := len(board)
	nums := make([]int, n*n+1)
	l2r, index := true, 1
	for i := n - 1; i >= 0; i-- {
		if l2r {
			for j := 0; j < n; j++ {
				nums[index] = board[i][j]
				index++
			}
		} else {
			for j := n - 1; j >= 0; j-- {
				nums[index] = board[i][j]
				index++
			}
		}

		l2r = !l2r
	}

	queue := []int{1}
	m := map[int]int{1: 0}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		step := m[cur]
		if cur == n*n {
			return step
		}

		for i := 1; i <= 6; i++ {
			next := cur + i
			if next > len(nums)-1 {
				continue
			}

			if nums[next] != -1 {
				next = nums[next]
			}
			if _, ok := m[next]; ok {
				continue
			}

			m[next] = step + 1
			queue = append(queue, next)
		}
	}

	return -1
}

// @lc code=end

