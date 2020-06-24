/*
 * @lc app=leetcode.cn id=401 lang=golang
 *
 * [401] 二进制手表
 */

// @lc code=start
// func readBinaryWatch(num int) []string {
// 	var res []string
// 	bits := make([]int, 60)
// 	for i := 0; i < 60; i++ {
// 		bits[i] = count(i)
// 	}

// 	for i := 0; i < 12; i++ {
// 		for j := 0; j < 60; j++ {
// 			if bits[i] + bits[j] == num {
// 				str := strconv.Itoa(i) + ":"
// 				if j < 10 {
// 					str += "0"
// 				}
// 				res = append(res, str + strconv.Itoa(j))
// 			}
// 		}
// 	}

// 	return res
// }

// func count(i int) int {
// 	res := 0
// 	for i > 0 {
// 		res++
// 		i = i & (i - 1)
// 	}
// 	return res
// }

func readBinaryWatch(num int) []string {
	var res []string
	bits := []int{1, 2, 4, 8, 1, 2, 4, 8, 16, 32}
	dfs(num, 0, 0, 0, bits, &res)
	return res
}

func dfs(num, hour, minute, index int, bits []int, res *[]string) {
	if hour > 11 || minute > 59 {
		return
	}

	if num == 0 {
		str := strconv.Itoa(hour) + ":"
		if minute < 10 {
			str += "0"
		}
		*res = append(*res, str + strconv.Itoa(minute))
		return
	}

	if index >= len(bits) {
		return
	}

	if num == 1 {
		if index < 4 {
			dfs(num - 1, hour + bits[index], minute, index + 1, bits, res)
		} else {
			dfs(num - 1, hour, minute + bits[index], index + 1, bits, res)
		}
	} else {
		for i := index + 1; i < len(bits); i++ {
			if index < 4 {
				dfs(num - 1, hour + bits[index], minute, index + 1, bits, res)
			} else {
				dfs(num - 1, hour, minute + bits[index], index + 1, bits, res)
			}
		}
	}
	dfs(num, hour, minute, index + 1, bits, res)

}
// @lc code=end

