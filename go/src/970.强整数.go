/*
 * @lc app=leetcode.cn id=970 lang=golang
 *
 * [970] 强整数
 */

// @lc code=start
// 2 1 10
func powerfulIntegers(x int, y int, bound int) []int {
	m := map[int]bool{}
	powers := make([][]int, 2)
	powers[0] = []int{1}
	powers[1] = []int{1}

	if x > 1 {
		xpower := 1
		for xpower * x <= bound {
			xpower *= x
			powers[0] = append(powers[0], xpower)
		}
	}

	if y > 1 {
		ypower := 1
		for ypower * y <= bound {
			ypower *= y
			powers[1] = append(powers[1], ypower)
		} 
	}


	xi, yi := 0, 0
	for {
		if yi == len(powers[1]) {
			yi = 0
			xi++
		}

		if xi == len(powers[0]) {
			break
		} 
		
		if powers[0][xi] + powers[1][yi] <= bound {
			m[powers[0][xi] + powers[1][yi]] = true
			yi++
			} else {
				xi++
				yi = 0
			}
	}


	res := []int{}
	for k, _ := range m {
		res = append(res, k)
	}
	return res
}
// @lc code=end

