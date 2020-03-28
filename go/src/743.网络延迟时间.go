/*
 * @lc app=leetcode.cn id=743 lang=golang
 *
 * [743] 网络延迟时间
 */

// @lc code=start

import (
	"math"
	"fmt"
)

type Edge struct {
	Sid int // 起始顶点编号
	Eid int // 终止顶点编号
	Weight int // 权重
}

func NewEdge(sid, eid, weight int) Edge {
	return Edge{
			Sid: sid,
			Eid: eid,
			Weight: weight,
	}
}

// 用于Dijkstra算法中
type Vertex struct {
	Id int // 顶点编号
	Dist int // 从起点到该点距离
}

func NewVertex(id, dist int) Vertex {
	return Vertex{
			Id: id,
			Dist: dist,
	}
}

type Graph struct {
	Count int // 顶点个数
	Adj [][]Edge // 邻接表
}

func NewGraph(count int) *Graph {
	return &Graph{    
			Count: count,
			Adj: make([][]Edge, count + 1), // 下标0 不使用
	}
}

func(this *Graph) AddEdge(a, b, weight int) { // 边a->b
	this.Adj[a] = append(this.Adj[a], NewEdge(a, b, weight))
}

func networkDelayTime(times [][]int, N int, K int) int {
	g := NewGraph(N)
	for i := 0; i < len(times); i++ {
		g.AddEdge(times[i][0], times[i][1], times[i][2])
	}

	vertexs := make([]Vertex, N + 1) // 记录从顶点K到各个顶点的距离
	for i := 1; i <= N; i++ {
		vertexs[i] = NewVertex(i, math.MaxInt32)
	}

	vertexs[K].Dist = 0
	queue := make([]Vertex, 1) // 最小堆
	queue = append(queue, vertexs[K])
	inqueue := make([]bool, N + 1) // 记录顶点是否已访问或在队列中
	inqueue[K] = true

	for len(queue) > 1 {
		minVertex := queue[1]
		queue[1] = queue[len(queue) - 1]
		queue = queue[:len(queue) - 1]
		heapify2(queue)

		for i := 0; i < len(g.Adj[minVertex.Id]); i++ {
			edge := g.Adj[minVertex.Id][i] 
			oldDist := vertexs[edge.Eid].Dist
			if minVertex.Dist + edge.Weight < oldDist {
					vertexs[edge.Eid].Dist = minVertex.Dist + edge.Weight
					if inqueue[edge.Eid] == false {
							//从下往上堆化
							queue = append(queue, vertexs[edge.Eid])
							heapify(queue)
							inqueue[edge.Eid] = true
					} else {
							update(queue, edge.Eid, vertexs[edge.Eid].Dist)
					}
			}
		}
	}

	max := 0
	for i := 1; i <= N; i++ {
		if vertexs[i].Dist == math.MaxInt32 {
				return -1 
		}
		if vertexs[i].Dist > max {
			max = vertexs[i].Dist
		}
	}
	return max
}

// 从下往上堆化
func heapify(queue []Vertex) {
	i := len(queue) - 1
	for i > 1 && queue[i].Dist < queue[i/2].Dist {
			queue[i], queue[i/2] = queue[i/2], queue[i]
			i /= 2
	}
}

// 从上往下堆化
func heapify2(queue []Vertex) {
	i := 1
	size := len(queue)
	for {
		min := i
		if 2*i < size && queue[i].Dist > queue[i*2].Dist {
			min = i * 2
		}
		if 2*i+1 < size && queue[min].Dist > queue[i*2+1].Dist {
			min = i * 2 + 1
		}


		if min == i {
			return
		}
		queue[i], queue[min] = queue[min], queue[i]
		i = min
	}
}

// 更新Id位id的顶点，因为每次更新后，该顶点的dist肯定是减少的，所以从下往上堆化即可
func update(queue []Vertex, id, newDist int) {
	var pos int
	for i := 1; i < len(queue); i++ {
			if queue[i].Id == id {
					pos = i
					break
			}
	}

	queue[pos].Dist = newDist
	for pos > 1 && queue[pos].Dist < queue[pos/2].Dist {
			queue[pos], queue[pos/2] = queue[pos/2], queue[pos]
			pos /= 2
	}
}
// @lc code=end

