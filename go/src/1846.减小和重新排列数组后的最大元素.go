/*
 * @lc app=leetcode.cn id=1846 lang=golang
 *
 * [1846] 减小和重新排列数组后的最大元素
 */

// @lc code=start
func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	n := len(arr)
	sort.Ints(arr)
	arr[0] = 1
	for i := 1; i < n; i++ {
		if arr[i]-arr[i-1] > 1 {
			arr[i] = arr[i-1] + 1
		}
	}

	return arr[n-1]
}

// @lc code=end

