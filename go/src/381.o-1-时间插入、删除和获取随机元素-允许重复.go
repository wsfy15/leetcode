type RandomizedCollection struct {
	index map[int][]int
	nums  []int
	count int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		index: make(map[int][]int),
		nums:  []int{},
		count: 0,
	}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	res := false
	if this.count < len(this.nums) {
		this.nums[this.count] = val
	} else {
		this.nums = append(this.nums, val)
	}

	if this.index[val] == nil {
		this.index[val] = []int{}
		res = true
	}
	this.index[val] = append(this.index[val], this.count)
	this.count++

	return res
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	if indexs, ok := this.index[val]; ok && len(indexs) > 0 {
		replaceVal := this.nums[this.count-1]
		this.nums[indexs[0]] = replaceVal
		for i := 0; i < len(this.index[replaceVal]); i++ {
			if this.index[replaceVal][i] == this.count-1 {
				this.index[replaceVal][i] = indexs[0]
				break
			}
		}

		this.index[val] = this.index[val][1:]
		this.count--
		return true
	}

	return false
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	return this.nums[rand.Intn(this.count)]
}
