package leetcode

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	MAX_LEVEL = 16 // 最高层数
)

// 以元素而非层数进行组织，每个元素（节点）的forwards切片长度与其层数一样，用于保存每一层的后继节点
// 跳表节点
type node struct {
	v        interface{} // 保存的值
	score    int         // 用于排序的分值 (存储对象的某个属性)
	level    int         // 层数
	forwards []*node     // 指向每层后继节点的指针
}

func newNode(v interface{}, score, level int) *node {
	return &node{v: v, score: score, level: level, forwards: make([]*node, level, level)}
}

// 跳表结构体
type SkipList struct {
	head   *node // 头节点，不存储实际数据，仅当作哨兵
	level  int   // 当前层数
	length int   // 最底层链表元素个数，即跳表长度
}

func NewSkipList() *SkipList {
	head := newNode(0, math.MinInt32, MAX_LEVEL)
	return &SkipList{head, 1, 0}
}

func (this *SkipList) Length() int {
	return this.length
}

func (this *SkipList) Level() int {
	return this.level
}

func (this *SkipList) Insert(v interface{}, score int) int {
	if nil == v {
		return 1
	}

	cur := this.head
	update := [MAX_LEVEL]*node{} // 记录每层的前驱节点

	for i := MAX_LEVEL - 1; i >= 0; i-- {
		// 找到每一层的前驱节点
		for nil != cur.forwards[i] {
			if cur.forwards[i].v == v {
				return 2
			}
			if cur.forwards[i].score > score {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}

		// 前面没找到前驱节点
		if nil == cur.forwards[i] {
			update[i] = cur
		}
	}

	// 通过随机算法获取该节点层数
	level := 1
	for i := 1; i < MAX_LEVEL; i++ {
		if rand.Int31()%7 == 1 {
			level++
		}
	}

	newNode := newNode(v, score, level)

	// 与每一层的前驱节点进行连接
	for i := 0; i < level; i++ {
		next := update[i].forwards[i] // 获取每一层前驱节点在该层的后继节点
		update[i].forwards[i] = newNode
		newNode.forwards[i] = next
	}

	if level > this.level {
		this.level = level
	}

	this.length++
	return 0
}

func (this *SkipList) Find(v interface{}, score int) *node {
	if nil == v || this.length == 0 {
		return nil
	}

	cur := this.head
	for i := this.level - 1; i >= 0; i-- {
		for nil != cur.forwards[i] {
			if cur.forwards[i].v == v && cur.forwards[i].score == score {
				return cur.forwards[i]
			}
			if cur.forwards[i].score > score {
				break
			}
			cur = cur.forwards[i]
		}
	}

	return nil
}

func (this *SkipList) Delete(v interface{}, score int) int {
	if nil == v {
		return 1
	}

	cur := this.head
	update := [MAX_LEVEL]*node{} // 记录前驱节点
	for i := this.level - 1; i >= 0; i-- {
		update[i] = this.head
		for nil != cur.forwards[i] {
			if cur.forwards[i].v == v && cur.forwards[i].score == score {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
	}

	cur = update[0].forwards[0] // 被删除节点
	for i := cur.level - 1; i >= 0; i-- {
		if update[i] == this.head && cur.forwards[i] == nil {
			this.level = i
		}

		if nil == update[i].forwards[i] {
			update[i].forwards[i] = nil
		} else {
			update[i].forwards[i] = update[i].forwards[i].forwards[i]
		}
	}

	this.length--

	return 0
}

func (this *SkipList) String() string {
	return fmt.Sprintf("level: %+v, length:%+v", this.level, this.length)
}
