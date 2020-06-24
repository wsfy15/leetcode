/*
 * @lc app=leetcode.cn id=705 lang=golang
 *
 * [705] 设计哈希集合
 */

// @lc code=start
type MyHashSet struct {
	bucketNum int
	buckets []*list.List
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {
	num := 1000
	hashSet := MyHashSet{
		bucketNum: num,
		buckets: make([]*list.List, num),
	}
	for i := 0; i < num; i++ {
		hashSet.buckets[i] = 	list.New()
	}
	return hashSet
}

func (this *MyHashSet) hash(key int) int {
	return key % this.bucketNum
}

func (this *MyHashSet) Add(key int)  {
	if this.Contains(key) {
		return
	}

	index := this.hash(key)
	this.buckets[index].PushBack(key)
}


func (this *MyHashSet) Remove(key int)  {
	index := this.hash(key)
	for e := this.buckets[index].Front(); e != nil; e = e.Next() {
		val := e.Value.(int)
		if val == key {
			this.buckets[index].Remove(e)
			return
		}
	}
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	index := this.hash(key)
	for e := this.buckets[index].Front(); e != nil; e = e.Next() {
		val := e.Value.(int)
		if val == key {
			return true
		}
	}
	return false
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
// @lc code=end

