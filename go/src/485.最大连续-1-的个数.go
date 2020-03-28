/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续1的个数
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
  max, cur := 0, 0
  for _, v := range nums {
    if v == 1 {
      cur++
    } else {
      if cur > max {
        max = cur
      }
      cur = 0
    }
  }
  if cur > max {
    return cur
  }
  return max
}
// @lc code=end

