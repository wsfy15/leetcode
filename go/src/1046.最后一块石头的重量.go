/*
 * @lc app=leetcode.cn id=1046 lang=golang
 *
 * [1046] 最后一块石头的重量
 */

// @lc code=start
func lastStoneWeight(stones []int) int {
	n := len(stones)
	if n == 1 {
		return stones[0]
	}
	if n == 2 {
		return abs(stones[1] - stones[0])
	}

	h := NewHeap(n)
	for _, num := range stones {
		h.Insert(num)
	}

	for h.Len() > 1 {
		a := h.Pop()
		b := h.Pop()
		if a - b > 0 {
			h.Insert(a - b)
		}
	}

	if h.Len() == 0 {
		return 0
	}
	return h.Pop()
}

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
	for i := this.Count; i/2 > 0 && this.Data[i] > this.Data[i/2]; {
		this.Data[i], this.Data[i/2] = this.Data[i/2], this.Data[i]
		i /= 2
	}
}

func(this *Heap) Len() int {
	return this.Count
}

func(this *Heap) Pop() int {
	tmp := this.Data[1]
	this.Data[1] = this.Data[this.Count]
	this.Count--
	this.heapify(1)
	return tmp
}

// i：对哪个下标的元素 从上到下堆化
func(this *Heap) heapify(i int) {
	for {
		maxPos := i	// 记录当前元素与其左右子节点中最大值 下标
		if i*2 <= this.Count && this.Data[i] < this.Data[i*2] {
			maxPos = i * 2
		}
		if i*2 + 1 <= this.Count && this.Data[maxPos] < this.Data[i*2 + 1] {
			maxPos = i * 2 + 1
		}
		if maxPos == i {
			return
		}
		
		this.Data[i], this.Data[maxPos] = this.Data[maxPos], this.Data[i]
		i = maxPos
	}
}

func abs(x int) int {
	y := x >> 31
	return (x ^ y) - y
}
// @lc code=end

