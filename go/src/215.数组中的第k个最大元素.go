/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
type Heap struct {
	Data []int	// 从下标1开始存储数据
	Count int	// 已存储元素个数
	Capacity int	// 堆中可以存储的最大元素个数
}

func NewHeap(capacity int) *Heap {
	return &Heap{
		Data: make([]int, capacity + 1),
		Count: 0,
		Capacity: capacity,
	}
}

func(this *Heap) Insert(data int) {
	if this.Count >= this.Capacity {
		return
	}
	
	this.Count++
	this.Data[this.Count] = data
	for i := this.Count; i/2 > 0 && this.Data[i] < this.Data[i/2]; {
		this.Data[i], this.Data[i/2] = this.Data[i/2], this.Data[i]
		i /= 2
	}
}

// 从上到下堆化
func(this *Heap) Heapify(i int) {
	for {
		maxPos := i	// 记录当前元素与其左右子节点中最大值 下标
		if i*2 <= this.Count && this.Data[i] > this.Data[i*2] {
			maxPos = i * 2
		}
		if i*2 + 1 <= this.Count && this.Data[maxPos] > this.Data[i*2 + 1] {
			maxPos = i * 2 + 1
		}
		if maxPos == i {
			return
		}
		
		this.Data[i], this.Data[maxPos] = this.Data[maxPos], this.Data[i]
		i = maxPos
	}
}

func findKthLargest(nums []int, k int) int {
	// 通过容量为k的最小堆
	// 或者快排
	heap := NewHeap(k)
	count := len(nums)
	for i := 0; i < k; i++ {
		heap.Insert(nums[i])
	}

	for i := k; i < count; i++ {
		if nums[i] > heap.Data[1] {
			heap.Data[1] = nums[i]
			heap.Heapify(1)
		}
	}

	return heap.Data[1]
}
// @lc code=end

