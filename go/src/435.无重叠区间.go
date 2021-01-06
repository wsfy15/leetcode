/*
 * @lc app=leetcode.cn id=435 lang=golang
 *
 * [435] 无重叠区间
 */

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {
	res := 0
	// 按起点排序，起点相同则按结尾排序
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	// 尽量保留终点小的
	for i := 0; i < len(intervals); i++ {
		j := i + 1
		for j < len(intervals) && intervals[i][1] > intervals[j][0] {
			if intervals[i][1] < intervals[j][1] {
				j++
				res++
			} else {
				res++
				break
			}
		}
		i = j - 1
	}

	return res
}

// @lc code=end

