/*
 * @lc app=leetcode.cn id=703 lang=golang
 *
 * [703] 数据流中的第 K 大元素
 */

// @lc code=start
type KthLargest struct {
	k int
	count int
	heap []int
}


func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{
		k: k,
		count: 0,
		heap: make([]int, k + 1),
	}

	for _, num := range nums {
		kl.insert(num)
	}
	return kl
}


func (this *KthLargest) Add(val int) int {
	this.insert(val)
	return this.heap[1]
}

func (this *KthLargest) insert(val int) {
	if this.count < this.k {
		this.count++
		this.heap[this.count] = val
		index := this.count
		for index > 1 && this.heap[index] < this.heap[index / 2] {
			this.heap[index], this.heap[index / 2] = this.heap[index / 2], this.heap[index]
			index /= 2 
		}
	} else if this.heap[1] < val {
		this.heap[1] = val
		index := 1
		for {
			pos := index
			if index * 2 <= this.k && this.heap[index] > this.heap[index * 2] {
				pos = index * 2
			}
			if index * 2 + 1 <= this.k && this.heap[pos] > this.heap[index * 2 + 1] {
				pos = index * 2 + 1
			}

			if pos == index {
				break
			}

			this.heap[index], this.heap[pos] = this.heap[pos], this.heap[index]
			index = pos
		}
	}
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
// @lc code=end

