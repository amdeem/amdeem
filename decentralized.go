package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

type Blockchain struct {
	blocks []*Block
	mu     sync.Mutex
}

type Item struct {
	ID          int
	Name        string
	Description string
	Price       float32
}

type Inventory struct {
	items []*Item
	mu    sync.Mutex
}

type Message struct {
	Type string
	Data []byte
}

var blockchain *Blockchain
var inventory *Inventory

func main() {
	// Initialize blockchain and inventory
	blockchain = &Blockchain{}
	inventory = &Inventory{}

	// Create and add genesis block
	genesisBlock := &Block{time.Now().Unix(), []byte("Genesis Block"), []byte{}, []byte{}, 0}
	genesisBlock.Hash = generateHash(genesisBlock)
	blockchain.blocks = append(blockchain.blocks, genesisBlock)

	// Serve HTTP endpoints
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/item", itemHandler)
	http.HandleFunc("/blockchain", blockchainHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Serve HTML page
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Decentralized Marketplace</title>
	</head>
	<body>
		<h1>Decentralized Marketplace</h1>
		<form action="/item" method="POST">
			<label for="name">Name:</label>
			<input type="text" name="name" required><br>
			<label for="description">Description:</label>
			<textarea name="description" required></textarea><br>
			<label for="price">Price:</label>
			<input type="number" name="price" step="0.01" min="0" required><br>
			<button type="submit">Post Item</button>
		</form>
	</body>
	</html>
	`
	fmt.Fprintf(w, html)
}

func itemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Parse form data
		name := r.FormValue("name")
		description := r.FormValue("description")
		priceStr := r.FormValue("price")
		price, err := strconv.ParseFloat(priceStr, 32)
		if err != nil {
			http.Error(w, "Invalid price", http.StatusBadRequest)
			return
		}

		// Add item to inventory
		item := &Item{len(inventory.items), name, description, float32(price)}
		inventory.mu.Lock()
		inventory.items = append(inventory.items, item)
		inventory.mu.Unlock()

		// Create message for item data
		itemData := &Message{"item", encodeGob(item)}
		itemDataBytes := encodeGob(itemData)

		// Create block with item data
		prevBlock := blockchain.blocks[len(blockchain.blocks)-1]
		newBlock := &Block{time.Now().Unix(), itemDataBytes, prevBlock.Hash, []byte{}, 0}
		proofOfWork(newBlock)
		blockchain.mu.Lock()
		blockchain.blocks = append(blockchain.blocks, newBlock)
		blockchain.mu.Unlock()

		fmt.Fprintf(w, "Item posted!")
	} else {
		// Invalid method
	
    
    
    
    ////---
    
    func itemHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        // Retrieve list of items from the blockchain and return to the client
        items := getItemsFromBlockchain()
        json.NewEncoder(w).Encode(items)
    case "POST":
        // Parse the request body to get the new item
        var newItem Item
        err := json.NewDecoder(r.Body).Decode(&newItem)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        // Create a new transaction to add the item to the blockchain
        tx := types.NewTransaction(nonce, common.HexToAddress("smart_contract_address_here"), nil, gasLimit, big.NewInt(0), data)

        // Sign the transaction using the user's private key
        signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Send the signed transaction to the blockchain
        err = client.SendTransaction(context.Background(), signedTx)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Return the new item to the client
        json.NewEncoder(w).Encode(newItem)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func getItemsFromBlockchain() []Item {
    // Connect to the blockchain using go-ethereum
    client, err := ethclient.Dial("http://localhost:8545")
    if err != nil {
        log.Fatal(err)
    }

    // Load the smart contract ABI
    abi, err := abi.JSON(strings.NewReader(string(contractABI)))
    if err != nil {
        log.Fatal(err)
    }

    // Get the contract address from the environment variable
    contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))

    // Get the contract instance using the ABI and address
    instance, err := NewMarketplace(contractAddress, client)
    if err != nil {
        log.Fatal(err)
    }

    // Get the number of items on the blockchain
    itemCount, err := instance.GetItemCount(nil)
    if err != nil {
        log.Fatal(err)
    }

    // Loop through each item and retrieve its details
    items := make([]Item, 0)
    for i := uint64(0); i < itemCount.Uint64(); i++ {
        name, description, price, seller, sold, err := instance.GetItem(nil, big.NewInt(int64(i)))
        if err != nil {
            log.Fatal(err)
        }

        // Convert the retrieved details into a struct and append to the list of items
        items = append(items, Item{
            Name:        name,
            Description: description,
            Price:       price,
            Seller:      seller,
            Sold:        sold,
        })
    }

    return items
}


    
    
    
    
    
    
    
    
    
