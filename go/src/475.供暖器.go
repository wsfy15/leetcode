/*
 * @lc app=leetcode.cn id=475 lang=golang
 *
 * [475] 供暖器
 */

// @lc code=start
func findRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)
	n, radius := len(heaters), 0
	for _, house := range houses {
		left, right := 0, n - 1
		for left < right { // 二分得到稍大于house的取暖器下标
			mid := left + (right - left) >> 1

			if heaters[mid] < house {
				left = mid + 1
			} else {
				right = mid
			}
		}

		distance := 0
		if heaters[left] < house { // 该加热器是最后一个加热器，所有加热器都在该house左边
			distance = house - heaters[left]
		} else if heaters[left] > house {
			distance = heaters[left] - house
			if left > 0 && house - heaters[left - 1] < distance {
				distance = house - heaters[left - 1]
			}
		}

		if distance > radius {
			radius = distance
		}
	}

	return radius
}
// @lc code=end

