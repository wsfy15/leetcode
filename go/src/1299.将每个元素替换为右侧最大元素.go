/*
 * @lc app=leetcode.cn id=1299 lang=golang
 *
 * [1299] 将每个元素替换为右侧最大元素
 */

// @lc code=start
func replaceElements(arr []int) []int {
	n := len(arr) - 1
	max := arr[n]
	arr[n] = -1
	n--
	for i := n; i >= 0; i-- {
		arr[i], max = max, fmax(max, arr[i])
	}
	return arr
}

func fmax(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res < v {
			res = v
		}
	}
	return res
}

// @lc code=end

