// Save this script as "generate_test_cases.go" within the "test" folder.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
)

const numTestCases = 10
const numVertices = 5

func main() {
	for i := 1; i <= numTestCases; i++ {
		testCase := generateTestCase()
		filename := fmt.Sprintf("test_case_%d.json", i)
		writeJSONToFile(testCase, filename)
	}
}

func generateTestCase() []string {
	vertices := []string{"a", "b", "c", "d", "e"}
	rand.Shuffle(len(vertices), func(i, j int) { vertices[i], vertices[j] = vertices[j], vertices[i] })

	numVerticesInTestCase := rand.Intn(numVertices) + 1 // Generate at least one vertex
	return vertices[:numVerticesInTestCase]
}

func writeJSONToFile(data interface{}, filename string) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
}
