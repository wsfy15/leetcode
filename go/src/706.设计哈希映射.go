/*
 * @lc app=leetcode.cn id=706 lang=golang
 *
 * [706] 设计哈希映射
 */

// @lc code=start
type pair struct {
	key int
	value int
}

type MyHashMap struct {
	bucketNum int
	buckets []*list.List
}


/** Initialize your data structure here. */
func Constructor() MyHashMap {
	num := 1000
	hashMap := MyHashMap{
		bucketNum: num,
		buckets: make([]*list.List, num),
	}
	for i := 0; i < num; i++ {
		hashMap.buckets[i] = 	list.New()
	}
	return hashMap
}

func (this *MyHashMap) hash(key int) int {
	return key % this.bucketNum
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
	index := this.hash(key)
	for e := this.buckets[index].Front(); e != nil; e = e.Next() {
		p := e.Value.(*pair)
		if p.key == key {
			p.value = value
			return
		}
	}

	this.buckets[index].PushBack(&pair{
		key: key,
		value: value,
	})
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	index := this.hash(key)
	for e := this.buckets[index].Front(); e != nil; e = e.Next() {
		p := e.Value.(*pair)
		if p.key == key {
			return p.value
		}
	}
	return -1
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
	index := this.hash(key)
	for e := this.buckets[index].Front(); e != nil; e = e.Next() {
		p := e.Value.(*pair)
		if p.key == key {
			this.buckets[index].Remove(e)
			return
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
// @lc code=end

