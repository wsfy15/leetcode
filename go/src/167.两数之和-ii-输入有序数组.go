/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
  counter := make(map[int]int)
  for i, v := range numbers {
    if index1, ok := counter[target-v]; ok {
      return []int{index1, i + 1}
    }
    counter[v] = i + 1
  }
  return nil
}
// @lc code=end

