package graph

import (
	"fmt"
)

// AMGraph this is undirected graph
type AMGraph struct {
	vertex []int //0~n
	edges  [][]int

	vertexNum, edgeNum int
}

func NewAMGraph(v int) *AMGraph {
	amg := new(AMGraph)
	amg.vertexNum = v
	amg.edges = make([][]int, v)
	for i := 0; i < v; i++ {
		amg.vertex = append(amg.vertex, i)
		amg.edges[i] = make([]int, v)
	}
	return amg
}

func (amg *AMGraph) AddEdge(x, y, w int) {
	if !amg.IsValidNode(x, y) {
		return
	}
	if amg.edges[x][y] == 0 {
		amg.edges[x][y] = w
		amg.edges[y][x] = w
		amg.edgeNum++
	}
}

func (amg *AMGraph) AddVertex() {
	n := amg.vertexNum
	amg.vertex = append(amg.vertex, n)
	for i := 0; i < n; i++ {
		amg.edges[i] = append(amg.edges[i], 0)
	}
	amg.edges = append(amg.edges, []int{})
	amg.edges[n] = make([]int, n+1)
	amg.vertexNum++
}

func (amg *AMGraph) DeleteVertex(x int) {
	if x <= 0 || x > amg.vertexNum {
		return
	}
	amg.vertex = append(amg.vertex[:x], amg.vertex[x+1:]...)
	for i := 0; i < amg.vertexNum; i++ {
		if i == x {
			continue
		}
		if amg.edges[x][i] != 0 {
			amg.edges[x][i] = 0
			amg.edges[i][x] = 0
		}
	}
	amg.vertexNum--
}
func (amg *AMGraph) SetEdgeWeight(x, y, w int) {
	if !amg.IsValidNode(x, y) {
		return
	}
	if amg.edges[x][y] == 0 {
		amg.edgeNum++
	}
	amg.edges[x][y] = w
}

func (amg *AMGraph) GetEdgeWeight(x, y int) (int, error) {
	if !amg.IsValidNode(x, y) {
		return 0, fmt.Errorf("%d and %d does not have edge", x, y)
	}
	return amg.edges[x][y], nil
}

func (amg *AMGraph) IsValidNode(x, y int) bool {
	if x == y || x < 0 || y < 0 || x >= amg.vertexNum || y >= amg.vertexNum {
		return false
	}
	return true
}

func (amg *AMGraph) Adjacent(x, y int) bool {
	if !amg.IsValidNode(x, y) {
		return false
	}
	return amg.edges[x][y] == 0
}

func (amg *AMGraph) Neighbor(x int) []int {
	if x <= 0 || x >= amg.vertexNum {
		return nil
	}
	return amg.edges[x]
}

func (amg *AMGraph) BFSTravel() {
	n := amg.vertexNum
	if n == 0 {
		return
	}
	var queue []int
	visited := make([]bool, n)
	visited[0] = true
	queue = append(queue, 0)
	ans := []int{0}
	for len(queue) != 0 {
		child := queue[0]
		queue = queue[1:]
		for i := 0; i < amg.vertexNum; i++ {
			if amg.edges[child][i] != 0 {
				if visited[i] {
					continue
				}
				queue = append(queue, i)
				visited[i] = true
				ans = append(ans, i)
			}
		}
	}
	fmt.Println(ans)
}

func (amg *AMGraph) DFSTravel() {
	n := amg.vertexNum
	if n == 0 {
		return
	}
	var stack []int
	visited := make([]bool, n)
	visited[0] = true
	stack = append(stack, 0)
	ans := []int{0}
	child := 0
	for true {
		for i := 0; i < amg.vertexNum; i++ {
			if amg.edges[child][i] != 0 {
				if !visited[i] {
					stack = append(stack, i)
					visited[i] = true
					child = i
					ans = append(ans, i)
				}
			}
		}
		if len(stack) == 0 {
			break
		}
		child = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	fmt.Println(ans)
}

func (amg *AMGraph) Display() [][]int {
	return amg.edges
}
