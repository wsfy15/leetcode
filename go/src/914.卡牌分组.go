/*
 * @lc app=leetcode.cn id=914 lang=golang
 *
 * [914] 卡牌分组
 */

// @lc code=start
func hasGroupsSizeX(deck []int) bool {
	arr := make([]int, 10000)
	for _, v := range deck {
		arr[v]++
	}

	g := -1
	for i := 0; i < 10000; i++ {
		if arr[i] > 0 {
			if g == -1 {
				g = arr[i]
			} else {
				g = gcd(g, arr[i])
			}
			if g < 2 {
				return false
			}
		}
	}
	return g >= 2
}

func hasGroupsSizeX2(deck []int) bool {
	count := make(map[int]int)
	for _, v := range deck {
		count[v]++
	}

	
	arr := []int{}
	for _, v := range count {
		arr = append(arr, v)
	}

	g := arr[0]
	for i := 1; i < len(arr); i++ {
		g = gcd(g, arr[i])
		if g < 2 {
			return false
		}
	}
	return g >= 2
}

func gcd(x, y int) int {
	if x < y {
		x, y = y, x
	}

	for x % y != 0 {
		x, y = y, x % y
	}
	return y
}

// @lc code=end

