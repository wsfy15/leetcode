/*
 * @lc app=leetcode.cn id=1287 lang=golang
 *
 * [1287] 有序数组中出现次数超过25%的元素
 */

// @lc code=start
func findSpecialInteger(arr []int) int {
	n := len(arr) / 4
	num, startIndex := arr[0], 0
	for i, v := range arr {
		if v != num {
			if i-startIndex > n {
				return num
			}

			num, startIndex = v, i
		}
	}
	return arr[0]
}

// @lc code=end

