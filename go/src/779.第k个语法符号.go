/*
 * @lc app=leetcode.cn id=779 lang=golang
 *
 * [779] 第K个语法符号
 */

// @lc code=start
func kthGrammar(N int, K int) int {
  // a := []int{0, 0, 1}
  // b := []int{0, 1, 0}
  count := 0
  for i := N - 2; i > 0; i-- {
    if K > (1 << i) {
      K -= 1 << i
      count++
    }
  }
  if count & 1 == 0 {
    return K - 1
    // return a[K]
  } else {
    return K & 1
    // return b[K]
  }
}
// @lc code=end

