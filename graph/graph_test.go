package graph

import (
	"fmt"
	"testing"
)

func TestAdListGraph_AddEdge(t *testing.T) {
	g := NewAdListGraph()
	n1, n2, n3, n4, n5 := Node{1}, Node{2}, Node{3}, Node{4}, Node{5}
	g.AddNode(&n1)
	g.AddNode(&n2)
	g.AddNode(&n3)
	g.AddNode(&n4)
	g.AddNode(&n5)

	g.AddEdge(&n1, &n2)
	g.AddEdge(&n1, &n5)
	g.AddEdge(&n2, &n3)
	g.AddEdge(&n2, &n4)
	g.AddEdge(&n2, &n5)
	g.AddEdge(&n3, &n4)
	g.AddEdge(&n4, &n5)
	fmt.Println(g.IsAdjacent(&n4, &n5))
	fmt.Println(g.IsAdjacent(&n4, &n1))
	g.Display()

	g.BFS(func(node *Node) {
		fmt.Printf("[Current Traverse Node]: %v\n", node)
	})

}
