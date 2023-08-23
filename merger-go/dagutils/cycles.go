package dagutils

import "fmt"

// In the dagutils package
func (dag *DAG) HasCycle() bool {
	visited := make(map[string]bool)
	recStack := make(map[string]bool)

	var hasCycleDFS func(string, []string) bool
	hasCycleDFS = func(vertex string, cycle []string) bool {
		fmt.Println("Visiting vertex:", vertex)
		if !visited[vertex] {
			visited[vertex] = true
			recStack[vertex] = true

			adjVertices := dag.Edges[vertex]
			for adjVertex := range adjVertices {
				if !visited[adjVertex] && hasCycleDFS(adjVertex, append(cycle, vertex)) {
					return true
				} else if recStack[adjVertex] {
					// Cycle detected, print the vertices in the cycle
					cycleStart := indexOf(cycle, adjVertex)
					if cycleStart != -1 {
						fmt.Println("Cycle detected:")
						fmt.Println("Vertices in the cycle:", cycle[cycleStart:])
					}
					return true
				}
			}
		}

		recStack[vertex] = false
		return false
	}

	for vertex := range dag.Vertices {
		if hasCycleDFS(vertex, nil) {
			return true
		}
	}
	return false
}

// Helper function to find the index of an element in a slice
func indexOf(slice []string, element string) int {
	for i, e := range slice {
		if e == element {
			return i
		}
	}
	return -1
}

// In the dagutils package
func (dag *DAG) BreakCycles() {
	// Initialize cycle index and cycle map
	cycleIndex := 0
	cycleMap := make(map[int][]string)

	// Helper function to perform DFS and mark cycles
	var markCyclesDFS func(string, []string)
	markCyclesDFS = func(vertex string, cycle []string) {
		visited := make(map[string]bool) // Define visited here
		visited[vertex] = true
		cycle = append(cycle, vertex)

		adjVertices := dag.Edges[vertex]
		for adjVertex := range adjVertices {
			if !visited[adjVertex] {
				markCyclesDFS(adjVertex, cycle)
			} else if visited[adjVertex] && cycleMap[cycleIndex] == nil {
				// Found a cycle
				cycleMap[cycleIndex] = cycle
				cycleIndex++
			}
		}
	}

	// Iterate over vertices and mark cycles
	for vertex := range dag.Vertices {
		if cycleMap[cycleIndex] == nil {
			markCyclesDFS(vertex, nil)
		}
	}

	// Create cycle vertices and update graph
	for _, cycle := range cycleMap {
		cycleVertex := fmt.Sprintf("Cycle%d", cycleIndex)
		dag.AddVertex(cycleVertex)
		for _, vertex := range cycle {
			dag.AddEdge(cycleVertex, vertex)
			dag.Vertices[vertex] = false // Mark original vertices as not included
		}
	}
}
