/*
 * @lc app=leetcode.cn id=165 lang=golang
 *
 * [165] 比较版本号
 */

// @lc code=start
func compareVersion(version1 string, version2 string) int {
	// 从头开始遍历每个区块
	len1, len2 := len(version1), len(version2)
	v1, v2, n1, n2 := 0, 0, 0, 0
	for n1 != len1 || n2 != len2 {
		v1, n1 = getNextChunk(version1, n1)
		v2, n2 = getNextChunk(version2, n2)
		if v1 != v2 {
			if v1 < v2 {
				return -1
			} else if v1 > v2 {
				return 1
			}
		}
	}

	return 0
}

func getNextChunk(version string, start int) (val int, end int) {
	if start == len(version) {
		return 0, start
	}

	end = start + 1
	for end < len(version) && version[end] != '.' {
		end++
	}
	val, _ = strconv.Atoi(version[start:end])
	if end < len(version) {
		// 说明是.
		end++ // 下一次开始的位置
	}
	return
}
// @lc code=end

