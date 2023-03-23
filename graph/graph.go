package graph

import (
	"fmt"
	"sync"
)

type Node struct {
	value int
}

type AdListGraph struct {
	nodes []*Node          // 节点集
	edges map[Node][]*Node // 邻接表表示的无向图
	lock  sync.RWMutex     // 保证线程安全
}

func NewAdListGraph() *AdListGraph {
	adListGraph := new(AdListGraph)
	adListGraph.edges = make(map[Node][]*Node)
	return adListGraph
}

func (adList *AdListGraph) AddNode(n *Node) {
	adList.lock.Lock()
	defer adList.lock.Unlock()
	adList.nodes = append(adList.nodes, n)
}

func (adList *AdListGraph) AddEdge(u, v *Node) {
	adList.lock.Lock()
	defer adList.lock.Unlock()
	adList.edges[*v] = append(adList.edges[*v], u)
	adList.edges[*u] = append(adList.edges[*u], v)
}

func (adList *AdListGraph) Display() {
	adList.lock.RLock()
	defer adList.lock.RUnlock()
	str := ""
	for _, node := range adList.nodes {
		str += node.String() + "->"
		nexts := adList.edges[*node]
		for _, next := range nexts {
			str += next.String() + " "
		}
		str += "\n"
	}
	fmt.Println(str)
}

func (adList *AdListGraph) IsAdjacent(u, v *Node) bool {
	list := adList.edges[*u]
	for _, node := range list {
		if node == v {
			return true
		}
	}
	return false
}

func (adList *AdListGraph) Neighbors(u *Node) {
	if adList.edges[*u] == nil {
		return
	}
	for _, edge := range adList.edges[*u] {
		fmt.Println(edge.value)
	}
}

func (g *AdListGraph) AdBFS() {
	g.lock.RLock()
	defer g.lock.RUnlock()

}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)

}

type NodeQueue struct {
	nodes []Node
	lock  sync.RWMutex
}

// 实现 BFS 遍历
func (g *AdListGraph) BFS(f func(node *Node)) {
	g.lock.RLock()
	defer g.lock.RUnlock()
	q := NewNodeQueue()
	if len(g.nodes) == 0 {
		return
	}
	head := g.nodes[0]
	q.Enqueue(*head)
	visited := make(map[*Node]bool)
	visited[head] = true
	for {
		if q.IsEmpty() {
			break
		}
		node := q.Dequeue()
		visited[node] = true
		for _, next := range g.edges[*node] {
			if visited[next] {
				continue
			}
			q.Enqueue(*next)
			visited[next] = true
		}
		if f != nil {
			f(node)
		}
	}
}

func (g *AdListGraph) DFS(f func(node *Node)) {

}

// 生成节点队列
func NewNodeQueue() *NodeQueue {
	q := NodeQueue{}
	q.lock.Lock()
	defer q.lock.Unlock()
	q.nodes = []Node{}
	return &q
}

// 入队列
func (q *NodeQueue) Enqueue(n Node) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.nodes = append(q.nodes, n)
}

// 出队列
func (q *NodeQueue) Dequeue() *Node {
	q.lock.Lock()
	defer q.lock.Unlock()
	node := q.nodes[0]
	q.nodes = q.nodes[1:]
	return &node
}

// 判空
func (q *NodeQueue) IsEmpty() bool {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return len(q.nodes) == 0
}
