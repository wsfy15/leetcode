/*
 * @lc app=leetcode.cn id=621 lang=golang
 *
 * [621] 任务调度器
 */

// @lc code=start
func leastInterval(tasks []byte, n int) int {
	counter := make([]int, 26)
	for _, ch := range tasks {
		counter[ch - 'A']++
	}	
	
	maxTime := max(counter...)
	maxTask := 0
	for _, num := range counter {
		if num == maxTime {
			maxTask++
		}
	}

	return max(len(tasks), (maxTime - 1) * (n + 1) + maxTask)
}

func max(a ...int) int {
	ans := a[0]
	for _, v := range a {
		if v > ans {
			ans = v
		}
	}
	return ans
}
// @lc code=end

