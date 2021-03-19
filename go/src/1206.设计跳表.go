/*
 * @lc app=leetcode.cn id=1206 lang=golang
 *
 * [1206] 设计跳表
 */

// @lc code=start
const MaxLevel = 16

type node struct {
	next [MaxLevel]*node
	val int
	level int // 当前节点最高层次
}

type Skiplist struct {
	head *node // 头节点，不存储实际数据
	level int  // 当前最大层数
	length int // 元素数量
}


func Constructor() Skiplist {
	head := &node{val: math.MinInt32, level: MaxLevel}
	return Skiplist{head, 1, 0}
}


func (this *Skiplist) Search(target int) bool {
	cur := this.head
	for i := this.level - 1; i >= 0; i-- {
		for cur.next[i] != nil {
			if cur.next[i].val == target {
				return true
			}
			if cur.next[i].val > target {
				break
			}

			cur = cur.next[i]
		}
	}

	return false
}


func (this *Skiplist) Add(num int)  {
	// 通过随机算法获取该节点层数
	level := 1
	for i := 1; i < MaxLevel; i++ {
		if rand.Int31()%7 == 1 {
			level++
		}
	}

	cur := this.head
	update := [MaxLevel]*node{} // 记录每层的前驱节点
	for i := level - 1; i >= 0; i-- {
		update[i] = this.head
		// 找到每一层的前驱节点
		for nil != cur {
			if cur.next[i] == nil || cur.next[i].val >= num {
				update[i] = cur
				break
			}
			cur = cur.next[i]
		}
	}

	n := node{val: num, level: level}
	for i := 0; i < level; i++ { // 在每一层插入当前节点
		n.next[i] = update[i].next[i]
		update[i].next[i] = &n
	}

	if level > this.level {
		this.level = level
	}
	this.length++
}


func (this *Skiplist) Erase(num int) bool {
	cur := this.head
	update := [MaxLevel]*node{} // 记录前驱节点
	for i := this.level - 1; i >= 0; i-- {
		for nil != cur.next[i] {
			if cur.next[i].val == num  {
				update[i] = cur
				break
			}
			if cur.next[i].val > num {
				break
			}
			cur = cur.next[i]
		}
	}

	if update[0] == nil {
		return false
	}

	cur = update[0].next[0] // 被删除节点
	for i := cur.level - 1; i >= 0; i-- {
		if update[i] == this.head && cur.next[i] == nil {
			this.level = i
		}

		update[i].next[i] = cur.next[i]
	}

	this.length--
	return true
}


/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
// @lc code=end
