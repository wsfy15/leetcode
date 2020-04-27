/*
 * @lc app=leetcode.cn id=384 lang=golang
 *
 * [384] 打乱数组
 */

// @lc code=start
type Solution struct {
	original []int
}


func Constructor(nums []int) Solution {
	original := make([]int, len(nums))
	copy(original, nums)
	return Solution{original}
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.original
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	size := len(this.original)
	array := make([]int, size)
	copy(array, this.original)

	for i := 0; i < size; i++ {
		randomIndex := i + rand.Intn(size - i)
		array[i], array[randomIndex] = array[randomIndex], array[i]
	}

	return array
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
// @lc code=end

