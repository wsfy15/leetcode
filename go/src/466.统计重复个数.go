/*
 * @lc app=leetcode.cn id=466 lang=golang
 *
 * [466] 统计重复个数
 */

// @lc code=start
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	index, s1cnt, s2cnt := 0, 0, 0
	len1, len2 := len(s1), len(s2)
	appear := map[int][2]int{} // key为s2的index，value为当前的s1cnt, s2cnt

	for s1cnt < n1 {
		s1cnt++
		for i := 0; i < len1; i++ {
			if s1[i] == s2[index] {
				index++
				if index == len2 {
					index = 0
					s2cnt++
				}
			}
		}

		// 用index-1或者index都可以
		// 只是为了标记s1扫描完后，之前是否也出现s2字符串判断到这个位置
		if v, ok := appear[index]; ok {
			// 遍历了(s1cnt - v[0])个s1，发现有(s2cnt - v[1])个s1
			s1num, s2num := s1cnt - v[0], s2cnt - v[1]
			left := n1 - s1cnt // 还剩下多少个s1
			loops := left / s1num // 还能重复多少次
			s2cnt += loops * s2num
			s1cnt += loops * s1num
			break
		} else {
			appear[index] = [2]int{s1cnt, s2cnt}
		}
	}

	for s1cnt < n1 {
		s1cnt++
		for i := 0; i < len1; i++ {
			if s1[i] == s2[index] {
				index++
				if index == len2 {
					index = 0
					s2cnt++
				}
			}
		}
	}

	return s2cnt / n2
}

func getMaxRepetitions2(s1 string, n1 int, s2 string, n2 int) int {
	count, p2 := 0, 0
	len1, len2 := len(s1), len(s2)
	for i := 0; i < n1; i++ {
		for j := 0; j < len1; j++ {
			if s1[j] == s2[p2] {
				p2++
				if p2 == len2 {
					count++
					p2 = 0
				}
			}
		}
	}

	return count/n2
}
// @lc code=end

