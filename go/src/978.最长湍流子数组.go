/*
 * @lc app=leetcode.cn id=978 lang=golang
 *
 * [978] 最长湍流子数组
 */

// @lc code=start
func maxTurbulenceSize(arr []int) int {
	n, res, up, down := len(arr), 1, 1, 1

	for i := 1; i < n; i++ {
		if arr[i] > arr[i-1] {
			up = down + 1
			down = 1
		} else if arr[i] < arr[i-1] {
			down = up + 1
			up = 1
		} else {
			down = 1
			up = 1
		}
		res = max(res, up, down)
	}

	return res
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res < v {
			res = v
		}
	}
	return res
}

// @lc code=end

