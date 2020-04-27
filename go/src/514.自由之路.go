/*
 * @lc app=leetcode.cn id=514 lang=golang
 *
 * [514] 自由之路
 */

// @lc code=start
func findRotateSteps(ring string, key string) int {
	pos := make(map[byte][]int) // ring中字符的位置
	for i := 0; i < len(ring); i++ {
		pos[ring[i]] = append(pos[ring[i]], i)
	}

	cache := make([][]int, len(ring))
	for i := 0; i < len(ring); i++ {
		cache[i] = make([]int, len(key))
	}

	return len(key) + dfs(ring, key, 0, 0, pos, cache)
}

func dfs(ring, key string, rindex, kindex int, pos map[byte][]int, cache [][]int) int {
	if kindex == len(key) {
		return 0
	}

	if cache[rindex][kindex] > 0 {
		return cache[rindex][kindex]
	}

	choiceNum, min := len(pos[key[kindex]]), math.MaxInt32
	for i := 0; i < choiceNum; i++ {
		p := pos[key[kindex]][i]
		cost := abs(rindex - p)
		if len(ring) - cost < cost {
			cost = len(ring) - cost
		}

		tmp := dfs(ring, key, p, kindex + 1, pos, cache)
		if tmp + cost  < min {
			min = tmp + cost
		}
	}
	cache[rindex][kindex] = min

	return cache[rindex][kindex]
}

func abs(n int) int {
	y := n >> 31
	return (n ^ y) - y
}
// @lc code=end

