/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
type Counter struct {
	Key int
	Val int
}

type Heap struct {
	Data []Counter	// 从下标1开始存储数据
	Count int	// 已存储元素个数
	Capacity int	// 堆中可以存储的最大元素个数
}

func NewHeap(capacity int) *Heap {
	return &Heap{
		Data: make([]Counter, capacity + 1),
		Count: 0,
		Capacity: capacity,
	}
}

func(this *Heap) Insert(key, val int) {
	if this.Count >= this.Capacity {
		return
	}
	
	this.Count++
	this.Data[this.Count].Key = key
	this.Data[this.Count].Val = val

	for i := this.Count; i/2 > 0 && this.Data[i].Val < this.Data[i/2].Val; {
		this.Data[i], this.Data[i/2] = this.Data[i/2], this.Data[i]
		i /= 2
	}
}

// 从上到下堆化
func(this *Heap) Heapify(i int) {
	for {
		maxPos := i	// 记录当前元素与其左右子节点中最大值 下标
		if i*2 <= this.Count && this.Data[i].Val > this.Data[i*2].Val {
			maxPos = i * 2
		}
		if i*2 + 1 <= this.Count && this.Data[maxPos].Val > this.Data[i*2 + 1].Val {
			maxPos = i * 2 + 1
		}
		if maxPos == i {
			return
		}
		
		this.Data[i], this.Data[maxPos] = this.Data[maxPos], this.Data[i]
		i = maxPos
	}
}

func topKFrequent(nums []int, k int) []int {
	heap := NewHeap(k)
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}

	count := 0
	for key, val := range m {
		count++
		if count > k {
			if val > heap.Data[1].Val {
				heap.Data[1] = Counter{key, val}
				heap.Heapify(1)
			}
		} else {
			heap.Insert(key, val)
		}
	}

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Data[i+1].Key 
	}
	return res
}
// @lc code=end

