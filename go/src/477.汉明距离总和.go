/*
 * @lc app=leetcode.cn id=477 lang=golang
 *
 * [477] 汉明距离总和
 */

// @lc code=start
func totalHammingDistance(nums []int) int {
	n, res := len(nums), 0
	bits := [31]int{} // 记录每一位1的个数
	for i := 0; i < n; i++ {
		index := 0
		for nums[i] > 0 {
			if nums[i]&1 == 1 {
				bits[index]++
			}
			index++
			nums[i] >>= 1
		}
	}

	for i := 0; i < 31; i++ {
		res += (n - bits[i]) * bits[i]
	}

	return res
}

// @lc code=end

