/*
 * @lc app=leetcode.cn id=1385 lang=golang
 *
 * [1385] 两个数组间的距离值
 */

// @lc code=start
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	res := 0
	for _, num := range arr1 {
		flag := true
		for _, num2 := range arr2 {
			if abs(num-num2) <= d {
				flag = false
				break
			}
		}
		if flag {
			res++
		}
	}

	return res
}

func abs(x int) int {
	y := x >> 31
	return (x ^ y) - y
}

// @lc code=end

