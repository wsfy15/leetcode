/*
 * @lc app=leetcode.cn id=561 lang=golang
 *
 * [561] 数组拆分 I
 */

// @lc code=start
func arrayPairSum(nums []int) int {
  sort.Ints(nums)
  res := 0
  for i := 0; i < len(nums); i += 2 {
      res += nums[i]
  }
  return res
}
// @lc code=end

