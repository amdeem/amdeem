package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// // Block represents a block in the blockchain.
type Block struct {
	Index        int
	Timestamp    int64
	Transactions []Transaction
	Validator    string
	PrevHash     string
	Hash         string
}

// Transaction represents a transaction in the blockchain.
type Transaction struct {
	Sender   string
	Receiver string
	Amount   float64
}

// Node represents a node in the blockchain network.
type Node struct {
	ID    string
	Stake float64
}

// Blockchain represents the blockchain.
type Blockchain struct {
	Chain        []Block
	Transactions []Transaction
	Nodes        []Node
}

// calculateHash calculates the SHA256 hash of a block.
func calculateHash(block Block) string {
	record := fmt.Sprintf("%d%d%v%s%s", block.Index, block.Timestamp, block.Transactions, block.Validator, block.PrevHash)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return fmt.Sprintf("%x", hashed)
}

// createBlock creates a new block in the blockchain.
func createBlock(transactions []Transaction, validator string, prevHash string) Block {
	block := Block{
		Index:        len(blockchain.Chain) + 1,
		Timestamp:    time.Now().UnixNano(),
		Transactions: transactions,
		Validator:    validator,
		PrevHash:     prevHash,
	}
	block.Hash = calculateHash(block)
	return block
}

// validateBlock validates a block by checking its hash and previous hash.
func validateBlock(block Block, prevBlock Block) bool {
	if block.PrevHash != prevBlock.Hash {
		return false
	}
	if calculateHash(block) != block.Hash {
		return false
	}
	return true
}

// addBlock adds a new block to the blockchain.
func addBlock(block Block) {
	blockchain.Chain = append(blockchain.Chain, block)
}

// validateChain validates the entire blockchain by checking the validity of each block.
func validateChain(chain []Block) bool {
	for i := 1; i < len(chain); i++ {
		if !validateBlock(chain[i], chain[i-1]) {
			return false
		}
	}
	return true
}

var blockchain Blockchain

func main() {
	// Create some nodes and add them to the blockchain.
	node1 := Node{ID: "node1", Stake: 10.0}
	node2 := Node{ID: "node2", Stake: 5.0}
	blockchain.Nodes = append(blockchain.Nodes, node1, node2)

	// Create some transactions and add them to the blockchain.
	tx1 := Transaction{Sender: "node1", Receiver: "node2", Amount: 2.5}
	blockchain.Transactions = append(blockchain.Transactions, tx1)

	// Select the validator based on their stake in the network.
	var validator string
	for _, node := range blockchain.Nodes {
		if node.Stake > 0 {
			validator = node.ID
			break
		}
	}

	// Create the new block and add it to the blockchain.
	prevHash := blockchain.Chain[len(blockchain.Chain)-1].Hash
	block := createBlock(blockchain.Transactions, validator, prevHash)
	addBlock(block)

	// Validate the entire blockchain.

	if validateChain(blockchain.Chain) {

		fmt.Println("Blockchain is not valid.")
	}
}
