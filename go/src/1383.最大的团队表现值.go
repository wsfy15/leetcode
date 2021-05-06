/*
 * @lc app=leetcode.cn id=1383 lang=golang
 *
 * [1383] 最大的团队表现值
 */

// @lc code=start
type item struct {
	speed      int
	efficiency int
}

type heap struct {
	k     int
	count int
	nums  []int
}

func NewHeap(k int) heap {
	return heap{
		k:     k,
		count: 0,
		nums:  make([]int, k+1),
	}
}

func (h *heap) push(num int) {
	h.count++
	h.nums[h.count] = num
	index := h.count
	for index/2 > 0 && h.nums[index] < h.nums[index/2] {
		h.nums[index], h.nums[index/2] = h.nums[index/2], h.nums[index]
		index /= 2
	}
}

func (h *heap) pop() int {
	res := h.nums[1]
	h.nums[1] = h.nums[h.count]
	h.count--
	index := 1
	for {
		pos := index
		if index*2 <= h.count && h.nums[index*2] < h.nums[index] {
			pos = index * 2
		}
		if index*2+1 <= h.count && h.nums[index*2+1] < h.nums[pos] {
			pos = index*2 + 1
		}

		if pos == index {
			break
		}

		h.nums[index], h.nums[pos] = h.nums[pos], h.nums[index]
		index = pos
	}

	return res
}

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	var mod int = 1e9 + 7
	items := make([]item, n)
	for i := 0; i < n; i++ {
		items[i] = item{speed[i], efficiency[i]}
	}

	// 根据效率降序排序，然后用大小为k的小顶堆存速度
	// 相当于固定效率值e，找到k-1个效率值不小于e的且速度总和最大的
	sort.Slice(items, func(i, j int) bool {
		return items[i].efficiency > items[j].efficiency
	})

	h := NewHeap(k) // 堆按速度排序
	res, sum := 0, 0
	for i := 0; i < n; i++ {
		if h.count >= k {
			sum -= h.pop()
		}

		sum += items[i].speed
		res = max(res, sum*items[i].efficiency)
		h.push(items[i].speed)
	}

	return res % mod
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

