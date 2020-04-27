/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
  m := make(map[int]int)
  for i, v := range nums {
    if i2, ok := m[target - v]; ok {
      return []int{i2, i}
    }
    m[v] = i
  }
  return nil
}
// @lc code=end

