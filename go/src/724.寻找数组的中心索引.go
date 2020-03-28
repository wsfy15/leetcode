/*
 * @lc app=leetcode.cn id=724 lang=golang
 *
 * [724] 寻找数组的中心索引
 */

// @lc code=start
func pivotIndex(nums []int) int {
  all := 0
  for _, v := range nums {
    all += v
  }
  
  cur := 0
  for i, v := range nums {
    if cur == all - v - cur {
      return i
    }
    cur += v
  }
  return -1
}
// @lc code=end

