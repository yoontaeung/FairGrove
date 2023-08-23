package txutils

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

const (
	minTransactionsPerSequence = 2
	hashLength                 = 5
)

type Transaction struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Value    string `json:"value"`
	GasPrice string `json:"gasPrice"`
	GasLimit string `json:"gasLimit"`
	Nonce    string `json:"nonce"`
	Data     string `json:"data"`
}

func calculateTransactionHash(transaction Transaction) string {
	transactionJSON, _ := json.Marshal(transaction)
	hash := sha256.Sum256([]byte(transactionJSON))
	hashString := hex.EncodeToString(hash[:])
	return hashString[:hashLength] // Truncate the hash to hashLength characters
}

func generateRandomTransaction() Transaction {
	from := "0x" + generateRandomHex(40) // Ethereum address format
	to := "0x" + generateRandomHex(40)
	value := generateRandomHex(18) // wei value
	gasPrice := generateRandomHex(10)
	gasLimit := generateRandomHex(6)
	nonce := generateRandomHex(8)
	data := generateRandomHex(64)

	return Transaction{
		From:     from,
		To:       to,
		Value:    value,
		GasPrice: gasPrice,
		GasLimit: gasLimit,
		Nonce:    nonce,
		Data:     data,
	}
}

func generateRandomHex(length int) string {
	bytes := make([]byte, length/2)
	rand.Read(bytes)
	return fmt.Sprintf("%x", bytes)
}

func generateTransactions(k int) ([]Transaction, []string) {
	transactions := make([]Transaction, k)
	transactionHashes := make([]string, k)
	for i := 0; i < k; i++ {
		transaction := generateRandomTransaction()
		transactions[i] = transaction
		transactionHashes[i] = calculateTransactionHash(transaction)
	}
	return transactions, transactionHashes
}

func writeTransactionsToFile(transactions []Transaction, filename string) error {
	data, err := json.MarshalIndent(transactions, "", "  ") // Indent with two spaces
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func writeTransactionHashesToFile(transactionHashes []string, filename string) error {
	hashesData, err := json.MarshalIndent(transactionHashes, "", "  ") // Indent with two spaces
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, hashesData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func TxGenerate(k int, n int) {
	// k: number of all transactions, n: number of sequences
	rand.Seed(time.Now().UnixNano())

	transactions, transactionHashes := generateTransactions(k)

	// Generate all transaction hashes
	var allTransactionHashes []string
	allTransactionHashes = append(allTransactionHashes, transactionHashes...)

	// Write all transaction hashes to a single file
	err := writeTransactionHashesToFile(allTransactionHashes, "all_transaction_hashes.json")
	if err != nil {
		fmt.Println("Error writing all transaction hashes:", err)
		return
	}

	maxTransactionsPerFile := 5

	// Calculate the minimum number of transactions per sequence
	minTransactionsPerSequence := minTransactionsPerSequence
	if k < minTransactionsPerSequence {
		minTransactionsPerSequence = k
	}

	usedIndices := 0
	for i := 1; i <= n; i++ {
		sharedIndices := rand.Perm(k) // Reset sharedIndices for each batch
		numTransactions := rand.Intn(maxTransactionsPerFile-minTransactionsPerSequence+1) + minTransactionsPerSequence

		if usedIndices+numTransactions > k {
			sharedIndices = rand.Perm(k)
			usedIndices = 0
		}

		selectedIndices := sharedIndices[usedIndices : usedIndices+numTransactions]
		usedIndices += numTransactions

		selectedTransactions := make([]Transaction, numTransactions)
		selectedTransactionHashes := make([]string, numTransactions)
		for j, index := range selectedIndices {
			selectedTransactions[j] = transactions[index]
			selectedTransactionHashes[j] = transactionHashes[index]
		}

		filename := fmt.Sprintf("transactions_%d.json", i)
		err := writeTransactionsToFile(selectedTransactions, filename)
		if err != nil {
			fmt.Println("Error writing transactions to file:", err)
			return
		}

		fmt.Printf("Transactions written to %s\n", filename)

		hashFilename := fmt.Sprintf("transactions_hash_%d.json", i)
		err = writeTransactionHashesToFile(selectedTransactionHashes, hashFilename)
		if err != nil {
			fmt.Println("Error writing transaction hashes to file:", err)
			return
		}

		fmt.Printf("Transaction hashes written to %s\n", hashFilename)
	}
}
