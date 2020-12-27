/*
 * @lc app=leetcode.cn id=882 lang=golang
 *
 * [882] 细分图中的可到达结点
 */

// @lc code=start
type Vertex struct {
	Id   int
	Dist int
}

func NewVertex(id, dist int) Vertex {
	return Vertex{
		Id:   id,
		Dist: dist,
	}
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
			min = i*2 + 1
		}

		if min == i {
			return
		}
		queue[i], queue[min] = queue[min], queue[i]
		i = min
	}
}

// 更新Id为id的顶点，因为每次更新后，该顶点的dist肯定是减少的，所以从下往上堆化即可
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

func reachableNodes(edges [][]int, M int, N int) int {
	if N == 1 || M == 0 {
		return 1
	}

	g := make([][]int, N) // 记录与每个顶点相关的边
	for i, edge := range edges {
		edges[i][2]++ // p1 p2 间距离要+1
		p1, p2 := edge[0], edge[1]
		g[p1] = append(g[p1], i)
		g[p2] = append(g[p2], i)
	}

	// 求点0到各点的距离
	dist := make([]int, N)
	for i := 1; i < N; i++ {
		dist[i] = 1e8
	}

	heap := []Vertex{NewVertex(0, 0), NewVertex(0, 0)} // 下标0的元素不使用
	inHeap := make([]bool, N)
	inHeap[0] = true
	for len(heap) > 1 {
		vertex := heap[1]
		heap[1] = heap[len(heap)-1]
		heap = heap[:len(heap)-1]
		heapify2(heap)

		for _, eIndex := range g[vertex.Id] {
			edge := edges[eIndex]
			another := edge[0]
			if another == vertex.Id {
				another = edge[1]
			}

			if dist[another] > dist[vertex.Id]+edge[2] {
				dist[another] = dist[vertex.Id] + edge[2]
				if inHeap[another] {
					update(heap, another, dist[another])
				} else {
					heap = append(heap, NewVertex(another, dist[another]))
					heapify(heap)
					inHeap[another] = true
				}
			}
		}
	}

	res := 0
	// 考虑原始图的顶点
	for i := 0; i < N; i++ {
		if dist[i] <= M {
			res++
		}
	}
	// 考虑新增的点
	for _, edge := range edges {
		p1, p2 := edge[0], edge[1]
		res += min(max(M-dist[p1], 0)+max(M-dist[p2], 0), edge[2]-1)
	}
	return res
}

func min(a ...int) int {
	ans := a[0]
	for _, v := range a {
		if v < ans {
			ans = v
		}
	}
	return ans
}

func max(a ...int) int {
	ans := a[0]
	for _, v := range a {
		if v > ans {
			ans = v
		}
	}
	return ans
}

// @lc code=end

