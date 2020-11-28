/*
 * @lc app=leetcode.cn id=164 lang=golang
 *
 * [164] 最大间距
 */

// @lc code=start
func maximumGap(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	// 将这些数放到不同的桶中，每个桶覆盖一个区间
	// 区间大小interval要使得要求的max gap不会出现在桶内
	// 而是在桶间，即当前桶的最小值减去上一个桶的最大值
	// 确保至少有一个桶没有元素，就能保证max gap不会出现在桶内
	// max(nums) 和 min(nums) 都不会参与计算，因为没有下一个桶的最小值和上一个桶的最大值
	// 因此考虑n-2个数，有n-1个桶就会保证至少有一个空桶
	// interval = (max - min) / (n - 1)

	minNum, maxNum := min(nums...), max(nums...)
	if minNum == maxNum {
		return 0
	}

	interval := int(math.Ceil(float64(maxNum-minNum) / float64(n-1)))
	bucketMin, bucketMax := make([]int, n-1), make([]int, n-1)
	for i := 0; i < n-1; i++ {
		bucketMin[i] = math.MaxInt32
		bucketMax[i] = -1
	}

	for _, num := range nums {
		if num == minNum || num == maxNum {
			continue
		}

		bucket := (num - minNum) / interval
		bucketMin[bucket] = min(num, bucketMin[bucket])
		bucketMax[bucket] = max(num, bucketMax[bucket])
	}

	res := 0
	previousMax := minNum
	for i := 0; i < n-1; i++ {
		// 这个桶没有数字
		if bucketMax[i] == -1 {
			continue
		}

		res = max(res, bucketMin[i]-previousMax)
		previousMax = bucketMax[i]
	}
	res = max(maxNum-previousMax, res)
	return res
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

func min(a ...int) int {
	ans := a[0]
	for _, v := range a {
		if v < ans {
			ans = v
		}
	}
	return ans
}

// @lc code=end

