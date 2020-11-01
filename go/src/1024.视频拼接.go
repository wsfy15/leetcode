/*
 * @lc app=leetcode.cn id=1024 lang=golang
 *
 * [1024] 视频拼接
 */

// @lc code=start
func videoStitching(clips [][]int, T int) int {
	if T == 0 {
		return 0
	}
	n := len(clips)
	if n == 0 {
		return -1
	}

	sort.Slice(clips, func(i, j int) bool {
		if clips[i][0] == clips[j][0] {
			return clips[i][1] > clips[j][1]
		}
		return clips[i][0] < clips[j][0]
	})

	if clips[0][0] != 0 {
		return -1
	}

	res, cur, right := 0, 0, 0
	for i := 0; i < n; i++ {
		// 已经覆盖了[0,cur]，考虑剩下的[cur, T]
		// 覆盖[0,cur]用了res个区间
		if clips[i][0] <= cur {
			// right是当前考虑到的区间能到达的最右位置
			right = max(right, clips[i][1])
		} else {
			if clips[i][0] > right {
				break
			}
			cur = right
			res++
			i--
		}

		if right >= T {
			break
		}
	}

	if cur < right {
		res++
		cur = right
	}

	if cur >= T {
		return res
	}
	return -1
}

func max(a ...int) int {
	ans := a[0]
	for _, v := range a {
		if v > ans {
			ans = v
		}
	}
	return ans
}

// @lc code=end

