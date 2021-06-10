/*
 * @lc app=leetcode.cn id=1738 lang=golang
 *
 * [1738] 找出第 K 大的异或坐标值
 */

// @lc code=start
type Heap struct {
	Data     []int // 从下标1开始存储数据
	Count    int   // 已存储元素个数
	Capacity int   // 堆中可以存储的最大元素个数
}

func NewHeap(capacity int) *Heap {
	return &Heap{
		Data:     make([]int, capacity+1),
		Count:    0,
		Capacity: capacity,
	}
}

func (this *Heap) Insert(data int) {
	if this.Count < this.Capacity {
		this.Count++
		this.Data[this.Count] = data
		for i := this.Count; i/2 > 0 && this.Data[i] < this.Data[i/2]; {
			this.Data[i], this.Data[i/2] = this.Data[i/2], this.Data[i]
			i /= 2
		}
	} else {
		if data > this.Data[1] {
			this.Data[1] = data
			this.heapify(1)
		}
	}
}

// i：对哪个下标的元素 从上到下堆化
func (this *Heap) heapify(i int) {
	for {
		maxPos := i // 记录当前元素与其左右子节点中最大值 下标
		if i*2 <= this.Count && this.Data[i] > this.Data[i*2] {
			maxPos = i * 2
		}
		if i*2+1 <= this.Count && this.Data[maxPos] > this.Data[i*2+1] {
			maxPos = i*2 + 1
		}
		if maxPos == i {
			return
		}

		this.Data[i], this.Data[maxPos] = this.Data[maxPos], this.Data[i]
		i = maxPos
	}
}

func kthLargestValue(matrix [][]int, k int) int {
	n, m := len(matrix), len(matrix[0])
	heap := NewHeap(k)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	dp[0][0] = matrix[0][0]
	heap.Insert(dp[0][0])
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] ^ matrix[i][0]
		heap.Insert(dp[i][0])
	}
	for j := 1; j < m; j++ {
		dp[0][j] = dp[0][j-1] ^ matrix[0][j]
		heap.Insert(dp[0][j])
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = dp[i-1][j] ^ dp[i][j-1] ^ dp[i-1][j-1] ^ matrix[i][j]
			heap.Insert(dp[i][j])
		}
	}

	return heap.Data[1]
}

// @lc code=end

