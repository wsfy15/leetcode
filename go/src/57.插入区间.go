/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	if n == 0 {
		return [][]int{newInterval}
	}
	
	var res [][]int
	start, end := newInterval[0], newInterval[1]

	if end < intervals[0][0] {
		return append([][]int{newInterval}, intervals...)
	}

	if start >  intervals[n-1][1] {
		return append(intervals, newInterval)
	}

	i := 0
	for ; i < n; i++ {
		if start <= intervals[i][1] {
			res = append(res, intervals[0:i]...)
			if end < intervals[i][0] {
				res = append(res, newInterval)
				break
			}

			if start > intervals[i][0] {
				start = intervals[i][0]
			} 
			if end < intervals[i][1] {
				end = intervals[i][1]
				i++
			} else {
				i++
				for i < n && end >= intervals[i][0]{
					i++
				}

				if end < intervals[i - 1][1] {
					end = intervals[i - 1][1]
				}
			}

			res = append(res, []int{start, end})
			break
		}
	}

	res = append(res, intervals[i:]...)
	return res
}
// @lc code=end

