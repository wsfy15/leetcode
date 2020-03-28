/*
 * @lc app=leetcode.cn id=703 lang=golang
 *
 * [703] 数据流中的第K大元素
 */

// @lc code=start
type KthLargest struct {
  heap []int // 存储k个元素的最小堆
  k int
  count int // 元素个数
}


func Constructor(k int, nums []int) KthLargest {
  heap := make([]int, k + 1)
  count := 0
  
  for i := 0; i < k && i < len(nums); i++ {
    heap[i + 1] = nums[i]
    count++
    for j := i + 1; j >= 2; j /= 2 {
      if heap[j] < heap[j / 2] {
        heap[j], heap[j / 2] = heap[j / 2], heap[j]
      } else {
        break
      }
    }
  }
  
  for i := k; i < len(nums); i++ {
    if nums[i] > heap[1] {
      heap[1] = nums[i]
      heapify(heap, k, 1)
    }
  }
  
  
  return  KthLargest{
    heap: heap, 
    k: k,
    count: count,
  }
}


func (this *KthLargest) Add(val int) int {
  if this.count < this.k {
    this.count++
    this.heap[this.count] = val
    for j := this.count; j >= 2; j /= 2 {
      if this.heap[j] < this.heap[j / 2] {
        this.heap[j], this.heap[j / 2] = this.heap[j / 2], this.heap[j]
      } else {
        break
      }
    }
  } else if val > this.heap[1] {
    this.heap[1] = val
    heapify(this.heap, this.k, 1)
  }
  
  return this.heap[1]
}

// 从上往下堆化
func heapify(heap []int, k, index int) {
  for index * 2 <= k {
    min := index
    if index*2 <= k && heap[index*2] < heap[index] {
      min = index * 2
    }
    if index*2 + 1 <= k && heap[index*2 + 1] < heap[min] {
      min = index * 2 + 1
    }
    
    if index == min {
      return
    }
    heap[min], heap[index] = heap[index], heap[min]
    index = min
  }
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
// @lc code=end

