/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 */

// @lc code=start
// 双向链表 + hash
// 重点是细心，链表节点的删除添加，指针的变化
type LRUCache struct {
	head *node // 最近访问的放在头节点
	tail *node
	count int // 当前数据数量
	capacity int
	hash map[int]*node
}

type node struct {
	key int
	value int
	prev *node
	next *node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		head: &node{},
		capacity: capacity,
		hash: make(map[int]*node),
	}
}


func (this *LRUCache) Get(key int) int {
	if node, ok := this.hash[key]; ok {
		if this.count > 1 {
			if this.tail == node {
				this.tail = node.prev
				this.tail.next = nil
			} else {
				node.next.prev = node.prev // 不是tail才有next
				node.prev.next = node.next
			}
			node.prev = this.head
			node.next = this.head.next
			this.head.next.prev = node
			this.head.next = node
		}
		return node.value
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	if this.Get(key) != -1 {
		this.head.next.value = value
		return
	}

	if this.count >= this.capacity {
		delete(this.hash, this.tail.key)
		this.tail = this.tail.prev
		this.tail.next = nil
		this.count--
	}

	this.count++
	node := &node{
		key: key,
		value: value,
		prev: this.head,
		next: this.head.next,
	}
	this.hash[key] = node

	if this.count == 1 {
		this.tail = node
	} else {
		this.head.next.prev = node
	}
	this.head.next = node
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
