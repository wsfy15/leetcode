/*
 * @lc app=leetcode.cn id=949 lang=golang
 *
 * [949] 给定数字能组成的最大时间
 */

// @lc code=start
func largestTimeFromDigits(A []int) string {
	// 4 个数字有 4 X 3 X 2 X 1 = 24 种组合
	res := make([]int, 4)
	res[0] = -1
	used := make([]bool, 4)
	dfs(A, used, []int{}, &res)

	if res[0] != -1 {
		return fmt.Sprintf("%02d:%02d", res[0] * 10 + res[1], res[2] * 10 + res[3])
	}
	return ""
}

func dfs(A []int, used []bool, cur []int, res *[]int) {
	allUse := true
	for i := 0; i < 4; i++ {
		if !used[i] {
			allUse = false
			cur = append(cur, A[i])
			used[i] = true
			dfs(A, used, cur, res)
			cur = cur[:len(cur) - 1]
			used[i] = false
		}
	}

	if allUse {
		hour := cur[0] * 10 + cur[1]
		minute := cur[2] * 10 + cur[3]
		if hour > 23 || minute > 59 {
			return
		}

		rhour := (*res)[0] * 10 + (*res)[1]
		rminute := (*res)[2] * 10 + (*res)[3]
		if hour > rhour || (hour == rhour && minute > rminute) {
			copy(*res, cur)
		}
	}
} 
// @lc code=end

