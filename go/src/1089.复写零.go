/*
 * @lc app=leetcode.cn id=1089 lang=golang
 *
 * [1089] 复写零
 */

// @lc code=start
func duplicateZeros(arr []int)  {
	n := len(arr)
	zeros := 0
	var index, v int
	for index, v = range arr {
		if v == 0 {
			zeros++
		}
		if index + zeros + 1 >= n {
			break
		}
	}

	duplicate := false
	if index + zeros + 1 > n {
		duplicate = true
	}
	for i := n - 1; i >= 0; i-- {
		if arr[index] == 0 {
			arr[i] = 0
			if duplicate {
				duplicate = false
				index--
			} else {
				duplicate = true
			}
		} else {
			arr[i] = arr[index]
			index--
		}
	}
}
// @lc code=end

