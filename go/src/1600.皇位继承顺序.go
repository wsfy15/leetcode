/*
 * @lc app=leetcode.cn id=1600 lang=golang
 *
 * [1600] 皇位继承顺序
 */

// @lc code=start
type node struct {
	name       string
	alive      bool
	successors []*node
}

type ThroneInheritance struct {
	root *node
	m    map[string]*node
}

func Constructor(kingName string) ThroneInheritance {
	ti := ThroneInheritance{
		m: make(map[string]*node),
	}
	king := ti.addPerson(kingName)
	ti.root = king
	return ti
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	child := this.addPerson(childName)
	parent := this.m[parentName]
	parent.successors = append(parent.successors, child)
}

func (this *ThroneInheritance) Death(name string) {
	this.m[name].alive = false
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
	return this.dfs(this.root, []string{})
}

func (this *ThroneInheritance) addPerson(name string) *node {
	child := &node{
		name:  name,
		alive: true,
	}
	this.m[name] = child
	return child
}

func (this *ThroneInheritance) dfs(cur *node, res []string) []string {
	if cur == nil {
		return res
	}

	if cur.alive {
		res = append(res, cur.name)
	}

	for _, child := range cur.successors {
		res = this.dfs(child, res)
	}

	return res
}

/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * obj := Constructor(kingName);
 * obj.Birth(parentName,childName);
 * obj.Death(name);
 * param_3 := obj.GetInheritanceOrder();
 */
// @lc code=end

