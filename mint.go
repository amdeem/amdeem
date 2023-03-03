type Block struct {
	Index     int
	Timestamp string
	Transactions []Transaction // new field for transactions
	Hash      string
	PrevHash  string
}

func NewBlock(index int, transactions []Transaction, prevHash string) *Block {
	block := &Block{
		Index:     index,
		Timestamp: time.Now().String(),
		Transactions: transactions, // set transactions field
		PrevHash:  prevHash,
	}
	block.Hash = calculateBlockHash(block)
	return block
}



type Blockchain struct {
	Blocks     []*Block
	Difficulty int
	PendingTx  []Transaction
	UserCoins  map[string]int // new field for user coin balances
}


type Transaction struct {
	Sender   string
	Receiver string
	Amount   int
	CoinReward int // new field for coin reward
}

func NewTransaction(sender string, receiver string, amount int, coinReward int) *Transaction {
	return &Transaction{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
		CoinReward: coinReward, // set coin reward field
	}
}

func (bc *Blockchain) AddTransaction(tx *Transaction) {
	bc.PendingTx = append(bc.PendingTx, *tx)
	if tx.Receiver != "system" {
		bc.updateUserCoins(tx.Receiver, tx.Amount + tx.CoinReward) // add coin reward to user's balance
	}
}

func (s *Server) HandlePostItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// create new transaction with coin reward
	tx := NewTransaction("system", item.Owner, 0, 1)

	// add transaction to pending transactions and update user's coin balance
	s.Blockchain.AddTransaction(tx)

	// add item to inventory
	s.Inventory = append(s.Inventory, item)

	w.WriteHeader(http.StatusCreated)
}












