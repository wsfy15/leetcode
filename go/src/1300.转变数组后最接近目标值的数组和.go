/*
 * @lc app=leetcode.cn id=1300 lang=golang
 *
 * [1300] 转变数组后最接近目标值的数组和
 */

// @lc code=start
func findBestValue(arr []int, target int) int {
	n := len(arr)
	sort.Ints(arr)
	tmp := divide(target, n)
	if tmp >= arr[n - 1] {
		return arr[n - 1]
	}
	if tmp <= arr[0] {
		return tmp
	}

	for i := 0; i < n - 1; i++ {
		target -= arr[i]
		tmp = divide(target, n - i - 1)
		if tmp >= arr[i] && tmp <= arr[i + 1] {
			// 因为四舍五入  XXX.5会变成(XXX+1)，而这种情况应该是XXX
			if abs((tmp - 1) * (n - i - 1) - target) == abs(tmp * (n - i - 1) - target) {
				return tmp - 1
			}
			return tmp
		}
	}
	return arr[n - 1]
}

func divide(x, y int) int {
	return int(math.Floor((float64(x) / float64(y)) + 0.5))
}

func abs(x int) int {
	y := x >> 31
	return (x ^ y) - y
}
// @lc code=end

