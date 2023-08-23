package dagutils

import (
	"fmt"
)

func (dag *DAG) TopologicalSort() ([]string, bool) {
	visited := make(map[string]bool)
	inProgress := make(map[string]bool)
	sortOrder := make([]string, 0)

	for vertex := range dag.Vertices {
		if !visited[vertex] {
			hasCycle := false
			sortOrder, hasCycle = dag.dfsTopologicalSort(vertex, visited, inProgress, sortOrder)
			if hasCycle {
				fmt.Println("Cycle detected during topological sort.")
				return nil, true
			}
		}
	}

	// Reverse the order to get the correct topological sort result
	ReverseOrder(sortOrder)
	fmt.Println("Topological sort completed successfully.")
	return sortOrder, false
}

func (dag *DAG) dfsTopologicalSort(vertex string, visited, inProgress map[string]bool, sortOrder []string) ([]string, bool) {
	visited[vertex] = false
	inProgress[vertex] = true

	for neighbor := range dag.Edges[vertex] {
		if inProgress[neighbor] {
			fmt.Println("Cycle detected:", neighbor, "is already in progress.")
			return nil, true // Cycle detected
		} else if !visited[neighbor] {
			var hasCycle bool
			sortOrder, hasCycle = dag.dfsTopologicalSort(neighbor, visited, inProgress, sortOrder)
			if hasCycle {
				fmt.Println("Cycle detected during DFS.")
				return nil, true
			}
		}
	}

	delete(inProgress, vertex)
	visited[vertex] = true
	sortOrder = append(sortOrder, vertex)
	return sortOrder, false
}

func ReverseOrder(slice []string) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func (dag *DAG) PrintSortOrder(sortOrder []string) {
	fmt.Println("Topological Sort Order:")
	for _, hash := range sortOrder {
		fmt.Println(hash)
	}
}
