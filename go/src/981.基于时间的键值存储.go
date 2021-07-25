/*
 * @lc app=leetcode.cn id=981 lang=golang
 *
 * [981] 基于时间的键值存储
 */

// @lc code=start
type node struct {
	value     string
	timestamp int
}

type TimeMap struct {
	m map[string][]node
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{
		m: make(map[string][]node),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	// 所有 TimeMap.set 操作中的时间戳 timestamps 都是严格递增的
	this.m[key] = append(this.m[key], node{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if nodes, ok := this.m[key]; ok {
		if timestamp < nodes[0].timestamp {
			return ""
		}

		l, r := 0, len(nodes)-1
		for l < r {
			m := l + (r-l+1)>>1
			if nodes[m].timestamp > timestamp {
				r = m - 1
			} else {
				l = m
			}
		}

		return nodes[l].value
	}

	return ""
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
// @lc code=end

