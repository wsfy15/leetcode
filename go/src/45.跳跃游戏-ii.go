/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	cur, count, farest := 0, 0, 1
	n := len(nums)

	for farest < n {
		tmp := farest
		for cur < tmp {
			next := cur + nums[cur] + 1
			if next >= n {
				return count + 1
			} else if next > farest {
				farest = next
			}
			cur++
		}
		count++
	}

	return count
}

func jump2(nums []int) int {
	n, count := len(nums), 0
	if n <= 1 {
		return 0
	}

	visited := make([]bool, n)
	queue, backupQueue := []int{0}, []int{}
	visited[0] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur + nums[cur] >= n - 1 {
			return count + 1
		}

		for i := cur + 1; i <= cur + nums[cur] && i < n; i++ {
			if !visited[i] {
				visited[i] = true
				backupQueue = append(backupQueue, i)
			}
		}

		if len(queue) == 0 && len(backupQueue) > 0 {
			queue, backupQueue = backupQueue, []int{}
			count++
		}
	}

	return count
}
// @lc code=end

