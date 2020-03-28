/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
  size := len(digits)
  over := 0
  res := make([]int, size)
  digits[size - 1]++
  for i := size - 1; i >= 0; i-- {
    res[i] = (digits[i] + over) % 10
    over = (digits[i] + over) / 10
  }
  
  if over == 1 {
    res = append([]int{1,}, res...)
  }
  
  return res
}
// @lc code=end

