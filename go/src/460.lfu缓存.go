/*
 * @lc app=leetcode.cn id=460 lang=golang
 *
 * [460] LFU缓存
 */

// @lc code=start

// 双哈希表
// 一个以frequency为key，一个以数据的key为key
type node struct {
	key int
	value int
	freq int
}

type LFUCache struct {
	capacity int
	count int
	minFreq int
	freq map[int]*list.List
	hash map[int]*list.Element
}


func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		freq: make(map[int]*list.List),
		hash: make(map[int]*list.Element),
		minFreq: 1,
	}
}


func (this *LFUCache) Get(key int) int {
	if ele, ok := this.hash[key]; ok {
		n := ele.Value.(*node)
		//fmt.Println("get before delete: ",this.freq[n.freq].Len())
		//this.freq[n.freq].Remove(ele)
		//fmt.Println("after delete: ",this.freq[n.freq].Len())
		for i := this.freq[n.freq].Front(); i != nil; i = i.Next() {
			n := i.Value.(*node)
			if key == n.key {
				this.freq[n.freq].Remove(i)
				break
			}
		}
		
		if this.minFreq == n.freq && this.freq[n.freq].Len() == 0 {
			this.minFreq++
		}

		n.freq++
		if l, ok := this.freq[n.freq]; ok {
			l.PushFront(n)
		} else {
			this.freq[n.freq] = list.New()
			this.freq[n.freq].PushFront(n)
		}

		return n.value
	}
	return -1
}


func (this *LFUCache) Put(key int, value int)  {
	if ele, ok := this.hash[key]; ok {
		n := ele.Value.(*node)
		//fmt.Println("put before delete: ",this.freq[n.freq].Len())
		//this.freq[n.freq].Remove(ele) // BUG！！！ 有可能删不掉
		//fmt.Println("after delete: ",this.freq[n.freq].Len())
		for i := this.freq[n.freq].Front(); i != nil; i = i.Next() {
			n := i.Value.(*node)
			if key == n.key {
				this.freq[n.freq].Remove(i)
				break
			}
		}

		if this.minFreq == n.freq && this.freq[n.freq].Len() == 0 {
			this.minFreq++
		}

		n.freq++
		n.value = value
		if l, ok := this.freq[n.freq]; ok {
			l.PushFront(n)
		} else {
			this.freq[n.freq] = list.New()
			this.freq[n.freq].PushFront(n)
		}
	} else {
		if this.capacity == 0 {
			return
		}

		if this.count == this.capacity {
			element := this.freq[this.minFreq].Back()
			n := element.Value.(*node)
			delete(this.hash, n.key)
			this.freq[this.minFreq].Remove(element)
			
			this.count--
		}

		this.count++
		n := &node{
			key: key,
			value: value,
			freq: 1,
		}

		if l, ok := this.freq[1]; ok {
			l.PushFront(n)
		} else {
			this.freq[1] = list.New()
			this.freq[1].PushFront(n)
		}

		this.hash[key] = this.freq[1].Front()
		this.minFreq = 1
	}
}

type node struct {
	key int
	val int
	count int
}

// 哈希表加双向链表
type LFUCache1 struct {
	l *list.List
	hash map[int]*list.Element
	count int
	capacity int
}


func Constructor1(capacity int) LFUCache1 {
	return LFUCache1{
		capacity: capacity,
		hash: make(map[int]*list.Element),
		l: list.New(),
	}
}


func (this *LFUCache1) Get(key int) int {
	if element, ok := this.hash[key]; ok {
		n := element.Value.(*node)
		n.count++
		for i := element.Next(); i != nil; i = i.Next() {
			n2 := i.Value.(*node)
			if n2.count > n.count {
				this.l.MoveBefore(element, i)
				return n.val
			}
		}
		this.l.MoveToBack(element)
		return n.val
	}
	return -1
}


func (this *LFUCache1) Put(key int, value int)  {
	if ele, ok := this.hash[key]; ok {
		n := ele.Value.(*node)
		n.val = value
		n.count++
		for i := ele.Next(); i != nil; i = i.Next() {
			n2 := i.Value.(*node)
			if n2.count > n.count {
				this.l.MoveBefore(ele, i)
				return
			}
		}
		this.l.MoveToBack(ele)
	} else {
		if this.count >= this.capacity {
			if this.count == 0 {
				return
			}
			
			this.count--
			head := this.l.Front()
			n := head.Value.(*node)
			delete(this.hash, n.key)
			this.l.Remove(head)
		}

		this.count++
		n := &node{
			key: key,
			val: value,
			count: 1,
		}

		for i := this.l.Front(); i != nil; i = i.Next() {
			n2 := i.Value.(*node)
			if n2.count > n.count {
				this.l.InsertBefore(n, i)
				this.hash[key] = i.Prev()
				return
			}
		}
		
		this.l.PushBack(n)
		this.hash[key] = this.l.Back()
	}
}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

