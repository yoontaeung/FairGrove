package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/dagutils"
)

func main() {
	originalDAG := dagutils.NewDAG()

	// Load transactions and construct the DAG
	for i := 1; ; i++ {
		filename := fmt.Sprintf("../test/test_case_%d.json", i)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			break
		}

		var transactionHashes []string
		err = json.Unmarshal(data, &transactionHashes)
		if err != nil {
			fmt.Println("Error unmarshaling transaction hashes:", err)
			return
		}

		for _, hash := range transactionHashes {
			originalDAG.AddVertex(hash)
		}

		for j := 1; j < len(transactionHashes); j++ {
			originalDAG.AddEdge(transactionHashes[j-1], transactionHashes[j])
		}
	}

	// Print or show the whole graph
	originalDAG.PrintGraph()
}
