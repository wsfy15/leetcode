/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int)  {
	n, left := len(nums), 0
	
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	} 

	for i := left; i < n; i++ {
		if nums[i] == 1 {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}
}
// @lc code=end

