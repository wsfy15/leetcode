/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int)  {
  // 反转两次 
  // init: 1,2,3,4,5,6,7
  // reverse first: 7,6,5,4,3,2,1
  // reverse first 3 elements: 5,6,7,4,3,2,1
  // reverse last 7 - 3 elements: 5,6,7,1,2,3,4
  size := len(nums)
  k %= size
  
  reverse(nums, 0, size - 1)
  reverse(nums, 0, k - 1)
  reverse(nums, k, size - 1)
}

func reverse(nums []int, start, end int) {
  for start < end {
    nums[start], nums[end] = nums[end], nums[start]
    start++
    end--
  }
}
// @lc code=end

