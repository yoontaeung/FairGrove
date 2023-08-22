package dagutils

import (
	"fmt"
)

type DAG struct {
	Vertices map[string]bool
	Edges    map[string]map[string]bool
}

func NewDAG() *DAG {
	return &DAG{
		Vertices: make(map[string]bool),
		Edges:    make(map[string]map[string]bool),
	}
}

func (dag *DAG) AddVertex(vertex string) {
	dag.Vertices[vertex] = true
}

func (dag *DAG) AddEdge(from, to string) {
	if dag.Edges[from] == nil {
		dag.Edges[from] = make(map[string]bool)
	}
	dag.Edges[from][to] = true
}

func (dag *DAG) PrintGraph() {
	fmt.Println("Vertices:")
	for vertex := range dag.Vertices {
		fmt.Println("-", vertex)
	}

	fmt.Println("Edges:")
	for from, toMap := range dag.Edges {
		for to := range toMap {
			fmt.Printf("- %s -> %s\n", from, to)
		}
	}
}
