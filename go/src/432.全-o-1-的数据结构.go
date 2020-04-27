/*
 * @lc app=leetcode.cn id=432 lang=golang
 *
 * [432] 全 O(1) 的数据结构
 */

// @lc code=start
// ["AllOne","inc","getMaxKey","getMinKey"]
type AllOne struct {
	Ref map[string] *list.Element // 一级索引，用来对元素进行定位，其value是一个kcnt结构
	LstRef map[int] *list.Element // 二级索引，用来对不同长度的链表进行定位，其value是一个链表。
	Lst *list.List // 二级链表，和二级索引LstRef相配对
}

type kcnt struct {
	k string
	cnt int
}


/** Initialize your data structure here. */
func Constructor() AllOne {
	rst:= AllOne{
			Ref: make(map[string] *list.Element),
			LstRef:make(map[int] *list.Element),
			Lst: list.New(),
	}
	elem:=rst.Lst.PushBack(list.New())
	rst.LstRef[0]=elem
	return rst
}


/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string)  {
	var cnt int
	if elem, ok := this.Ref[key]; ok {
			// 获取计数cnt
			cnt = elem.Value.(*kcnt).cnt
			// 从old链中删除
			this.LstRef[cnt].Value.(*list.List).Remove(elem)
	}else{
			// 计数为0
			cnt = 0 
	}
	// 如果cnt链没有，则新建
	if _,ok := this.LstRef[cnt+1]; !ok {
			elem:=this.Lst.InsertAfter(list.New(),this.LstRef[cnt]) // 这里需要初始化一个cnt为0的一个空链表         
			this.LstRef[cnt+1]=elem
	}
	// 应该先加后删除
	if this.LstRef[cnt].Value.(*list.List).Len()==0 && cnt > 0 {
			this.Lst.Remove(this.LstRef[cnt])
			delete(this.LstRef,cnt)
	}
	// cnt链中添加元素，返回*list.Element
	elem := this.LstRef[cnt+1].Value.(*list.List).PushBack(&kcnt{k:key,cnt:cnt+1})
	// 更新Ref[key]
	this.Ref[key]=elem
}


/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string)  {
	var cnt int
	if elem,ok := this.Ref[key]; !ok {
			return
	}else {
			cnt = elem.Value.(*kcnt).cnt
			// 从老链中删除
			this.LstRef[cnt].Value.(*list.List).Remove(elem)
			delete(this.Ref,key)
	}

	// 如果新链没有，则创建
	if _,ok:=this.LstRef[cnt-1];!ok {
			elem:=this.Lst.InsertBefore(list.New(),this.LstRef[cnt])
			this.LstRef[cnt-1]=elem
	}
	if this.LstRef[cnt].Value.(*list.List).Len()==0{
			this.Lst.Remove(this.LstRef[cnt])
			delete(this.LstRef,cnt)
	}
	// 将key，cnt加入新链
	if cnt==1{
			return
	}
	elem := this.LstRef[cnt-1].Value.(*list.List).PushBack(&kcnt{k:key,cnt:cnt-1})
	// 更新Ref[key]
	this.Ref[key]=elem
}


/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMinKey() string {
	if this.Lst.Len()<2{
			return ""
	}
	lst:=this.Lst.Front().Next().Value.(*list.List)
	return lst.Front().Value.(*kcnt).k
}


/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMaxKey() string {
	if this.Lst.Len()<2{
			return ""
	}
	lst:=this.Lst.Back().Value.(*list.List)
	return lst.Front().Value.(*kcnt).k
}



// type AllOne struct {
// 	k2v map[string]int
// 	freq map[int]*list.List
// 	keyLoc map[string]*list.Element
// 	min int
// 	max int
// 	count int
// }


// /** Initialize your data structure here. */
// func Constructor() AllOne {
// 	return AllOne{
// 		k2v: make(map[string]int),
// 		freq: make(map[int]*list.List),
// 		keyLoc: make(map[string]*list.Element),
// 		min: 1,
// 	}
// }


// /** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
// func (this *AllOne) Inc(key string)  {
// 	oldValue := 0
// 	if v, ok := this.k2v[key]; ok {
// 		oldValue = v
// 	} else {
// 		this.count++
// 	}

// 	this.update(key, oldValue, oldValue + 1)
// }


// /** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
// func (this *AllOne) Dec(key string)  {
// 	oldValue := 0
// 	if v, ok := this.k2v[key]; ok {
// 		oldValue = v
// 	} else {
// 		return
// 	}
	
// 	this.update(key, oldValue, oldValue - 1)
// }


// /** Returns one of the keys with maximal value. */
// func (this *AllOne) GetMaxKey() string {
// 	if  this.count == 0 {
// 		return ""
// 	}

// 	node := this.freq[this.max].Front()
// 	return node.Value.(string)
// }


// /** Returns one of the keys with Minimal value. */
// func (this *AllOne) GetMinKey() string {
// 	if this.count == 0 {
// 		return ""
// 	}

// 	node := this.freq[this.min].Front()
// 	return node.Value.(string)
// }

// func (this *AllOne) update(key string, oldValue, newValue int) {
// 	if newValue > this.max {
// 		this.max = newValue
// 	} else if newValue > 0 && newValue < this.min {
// 		this.min = newValue
// 	}

// 	this.k2v[key] = newValue
// 	if oldValue > 0 {
// 		this.freq[oldValue].Remove(this.keyLoc[key])
// 	}

// 	if newValue == 0 {
// 		delete(this.k2v, key)
// 		delete(this.keyLoc, key)
// 		this.count--

// 		// 可能当前最小值需要更新
// 		this.min = 1
// 		for i := 1; i <= this.max; i++ {
// 			if this.freq[i].Len() > 0 {
// 				this.min = i
// 				return
// 			}
// 		}
// 		return
// 	}

// 	if _, ok := this.freq[newValue]; !ok {
// 		this.freq[newValue] = list.New()
// 	}

// 	this.freq[newValue].PushFront(key)
// 	this.keyLoc[key] = this.freq[newValue].Front()

// 	if oldValue == this.max && newValue < oldValue && newValue > 0 {
// 		// 减少的时候，可能最大值会减小
// 		if this.freq[this.max].Len() == 0 {
// 			this.max--
// 		}
// 	} else if oldValue == this.min && newValue > oldValue {
// 		//	增加的时候，可能最小值也跟着增加
// 		if this.freq[this.min].Len() == 0 {
// 			this.min++
// 		}
// 	}
// }


/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
// @lc code=end

