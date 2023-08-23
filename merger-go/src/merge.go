package main

import (
	"main/txutils"
)

func main() {
	numSequencers := 3
	totalTransactions := 5
	txutils.TxGenerate(totalTransactions, numSequencers)

	// fmt.Println("Local Sequences:")
	// for i, localSequence := range localSequences {
	// 	fmt.Printf("Sequencer %d: %v\n", i, localSequence)
	// }
}
