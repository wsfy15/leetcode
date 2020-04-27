/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] <= intervals[j][1]
		}
		return intervals[i][0] <= intervals[j][0]
	})

	res := [][]int{}
	n := len(intervals)
	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < n; i++ {
		if intervals[i][0] <= end {
			if intervals[i][1] > end {
				end = intervals[i][1]
			}
		} else {
			res = append(res, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		}
	}
	res = append(res, []int{start, end})

	return res
}
// @lc code=end

