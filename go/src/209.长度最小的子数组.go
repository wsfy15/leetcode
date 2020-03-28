/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(s int, nums []int) int {
  min, start, acc := len(nums) + 1, 0, 0
  for i, v := range nums {
    acc += v
    for acc >= s && acc - nums[start] >= s {
      acc -= nums[start]
      start++
    }
    if acc >= s {
      if i - start + 1 < min {
        min =  i - start + 1
      }
    }
  }
  if min == len(nums) + 1 {
    return 0
  }
  return min
}
// @lc code=end

