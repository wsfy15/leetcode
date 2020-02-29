/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 */

// @lc code=start
type MedianFinder struct {
	BigHeapData []int	// 大元素顶堆存储前半部分 + 1个元素，堆顶元素即中位数
	BHCount int	
	SmallHeapData []int	// 小元素顶堆存储后半部分
	SHCount int 
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		BigHeapData: make([]int, 1),
		SmallHeapData: make([]int, 1),
	}
}


func (this *MedianFinder) AddNum(num int)  {
	if this.BHCount == 0 && this.SHCount == 0 {
		this.BigHeapData = append(this.BigHeapData, num)
		this.BHCount++
		return
	}

	if num < this.BigHeapData[1] {
		this.BigHeapData = append(this.BigHeapData, num)
		this.BHCount++
		for i := this.BHCount; i > 0; i /= 2 {
			if this.BigHeapData[i] > this.BigHeapData[i/2] {
				this.BigHeapData[i], this.BigHeapData[i/2] = this.BigHeapData[i/2], this.BigHeapData[i]
			} else {
				break
			}
		}
	} else {
		this.SmallHeapData = append(this.SmallHeapData, num)
		this.SHCount++
		for i := this.SHCount; i > 0; i /= 2 {
			if this.SmallHeapData[i] < this.SmallHeapData[i/2] {
				this.SmallHeapData[i], this.SmallHeapData[i/2] = this.SmallHeapData[i/2], this.SmallHeapData[i]
			} else {
				break
			}
		}
	}

	if this.BHCount > this.SHCount + 1 {
		this.SmallHeapData = append(this.SmallHeapData, this.BigHeapData[1])
		this.SHCount++
		for i := this.SHCount; i / 2 > 0; i /= 2 {
			if this.SmallHeapData[i] < this.SmallHeapData[i/2] {
				this.SmallHeapData[i], this.SmallHeapData[i/2] = this.SmallHeapData[i/2], this.SmallHeapData[i]
			} else {
				break
			}
		}

		this.BigHeapData[1] = this.BigHeapData[this.BHCount]
		this.BigHeapData = this.BigHeapData[:this.BHCount]
		this.BHCount--
		for i := 1; i < this.BHCount; {
			maxPos := i
			if i*2 <= this.BHCount && this.BigHeapData[i] < this.BigHeapData[2*i] {
				maxPos = 2*i
			}
			if i*2 + 1 <= this.BHCount && this.BigHeapData[maxPos] < this.BigHeapData[2*i+1] {
				maxPos = 2*i+1
			}

			if maxPos == i {
				break
			}
			this.BigHeapData[i], this.BigHeapData[maxPos] = this.BigHeapData[maxPos], this.BigHeapData[i]
			i = maxPos
		}
	} else if this.BHCount < this.SHCount {
		this.BigHeapData = append(this.BigHeapData, this.SmallHeapData[1])
		this.BHCount++
		for i := this.BHCount; i / 2 > 0; i /= 2 {
			if this.BigHeapData[i] > this.BigHeapData[i/2] {
				this.BigHeapData[i], this.BigHeapData[i/2] = this.BigHeapData[i/2], this.BigHeapData[i]
			} else {
				break
			}
		}

		this.SmallHeapData[1] = this.SmallHeapData[this.SHCount]
		this.SmallHeapData = this.SmallHeapData[:this.SHCount]
		this.SHCount--
		for i := 1; i < this.SHCount; {
			maxPos := i
			if i*2 <= this.SHCount && this.SmallHeapData[i] > this.SmallHeapData[2*i] {
				maxPos = 2*i
			}
			if i*2 + 1 <= this.SHCount && this.SmallHeapData[maxPos] > this.SmallHeapData[2*i+1] {
				maxPos = 2*i+1
			}

			if maxPos == i {
				break
			}
			this.SmallHeapData[i], this.SmallHeapData[maxPos] = this.SmallHeapData[maxPos], this.SmallHeapData[i]
			i = maxPos
		}
	}
}


func (this *MedianFinder) FindMedian() float64 {
	// fmt.Printf("%+v\n", this) // 加上这行会超时
	if (this.BHCount + this.SHCount) % 2 == 1 {
		return float64(this.BigHeapData[1])
	} else {
		return float64(this.BigHeapData[1] + this.SmallHeapData[1]) / 2.0
	}
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end

