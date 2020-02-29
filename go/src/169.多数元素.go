/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	// 排序后下标为 n/2 的元素即为众数
	// or Boyer-Moore 算法
	var count int = 0
	var e int
	for _, num := range nums {
		if count == 0 {
			e = num
		}
		if num == e {
			count += 1
		} else {
			count -= 1
		}
	}

	return e
}

// @lc code=end
