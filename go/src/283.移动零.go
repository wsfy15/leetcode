/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int)  {
  size := len(nums)
  if size == 0 {
    return
  }
  
  j := 0
  for _, v := range nums {
    if v != 0 {
      nums[j] = v
      j++
    }
  }
  for j < size {
    nums[j] = 0
    j++
  }
}
// @lc code=end

