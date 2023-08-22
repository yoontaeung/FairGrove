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
	// Check if the vertices exist
	if dag.Vertices[from] && dag.Vertices[to] {
		// Check if the edge already exists
		if dag.Edges[from][to] {
			// Edge from 'from' to 'to' already exists, no need to add
			return
		}

		// Check if the opposite edge exists
		if dag.Edges[to][from] {
			// Edge from 'to' to 'from' exists, remove it
			delete(dag.Edges[to], from)
			return
		}

		// Add the new edge from 'from' to 'to'
		if dag.Edges[from] == nil {
			dag.Edges[from] = make(map[string]bool)
		}
		dag.Edges[from][to] = true
	}
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
