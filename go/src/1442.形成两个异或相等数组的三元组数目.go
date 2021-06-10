/*
 * @lc app=leetcode.cn id=1442 lang=golang
 *
 * [1442] 形成两个异或相等数组的三元组数目
 */

// @lc code=start
func countTriplets(arr []int) int {
	n, res := len(arr), 0
	for i := 0; i < n; i++ {
		cur := arr[i]
		for j := i + 1; j < n; j++ {
			cur ^= arr[j]
			if cur == 0 {
				res += j - i
			}
		}
	}

	return res
}

// @lc code=end

