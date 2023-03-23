package graph

import (
	"fmt"
	"testing"
)

func TestAMGraph(t *testing.T) {
	g := NewAMGraph(4)
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 2, 1)
	g.AddEdge(1, 3, 1)
	fmt.Println(g.edges)
	g.DFSTravel()
}
