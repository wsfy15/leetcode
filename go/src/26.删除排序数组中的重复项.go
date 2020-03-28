/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
  size := len(nums)
  if size == 0 {
    return 0
  }
  
  last, j := nums[0] + 1, 0
  for _, v := range nums {
    if last != v {
      nums[j] = v
      last = v
      j++
    }
  }
  return j
}
// @lc code=end

