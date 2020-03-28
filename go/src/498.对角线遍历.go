/*
 * @lc app=leetcode.cn id=498 lang=golang
 *
 * [498] 对角线遍历
 */

// @lc code=start
func findDiagonalOrder(matrix [][]int) []int {
  rows := len(matrix)
  if rows == 0 {
    return nil
  }
  cols := len(matrix[0])
  
  // 运动规律：右上 -> 左下 -> 右下 -> 左下 -> ...... 右上时到达边界，则变成向右；坐下同理
  dx := -1 // []int{-1, 1}
  dy := 1  // []int{1, -1}
  x := 0
  y := 0
  // k := 0
  res := make([]int, rows * cols)
  count := 0
  for count < rows * cols {
    res[count] = matrix[x][y]
    count++
    
    // update x y
    x += dx
    y += dy
    if x == -1 || y == -1 || x == rows || y == cols {
      // k = (k + 1) % 2
      dx *= -1
      dy *= -1
      if x == -1 {
        x = 0
        if y == cols { // 右上顶点
          x++
          y--
        }
      } else if y == -1 {
        y++
        if x == rows { // 左下顶点
          x--
          y++
        }
      } else if x == rows {
        x--
        y += 2
      } else if y == cols {
        y--
        x += 2
      }
    }
  }
  
  return res
}
// @lc code=end

