/*
 * @lc app=leetcode.cn id=747 lang=golang
 *
 * [747] 至少是其他数字两倍的最大数
 */

// @lc code=start
func dominantIndex(nums []int) int {
  maxIndex := 0
  for i, v := range nums {
    if v > nums[maxIndex] {
      maxIndex = i
    }
  }
  for  i, v := range nums {
    if i == maxIndex {
      continue
    }
    if v * 2 > nums[maxIndex] {
      return -1
    }
  }
  return maxIndex
}
// @lc code=end

