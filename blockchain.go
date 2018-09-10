package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"github.com/boltdb/bolt"
	"strconv"
	"os"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

// Blockchain keeps a sequence of Blocks
//区块链，保存了最近的区块的hash以及数据库对象
type Blockchain struct {
	tip []byte
	db  *bolt.DB
}

// BlockchainIterator is used to iterate over blockchain blocks
//迭代器，用于循环区块
type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

// AddBlock saves provided data as a block in the blockchain
//增加一个区块
func (bc *Blockchain) AddBlock() {
	var lastHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock([]*Transaction{},lastHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		bc.tip = newBlock.Hash
		//fmt.Printf("Prev. hash: %x\n", newBlock.PrevBlockHash)
		//fmt.Printf("Data: %s\n", newBlock.MerkleRoot)
		//fmt.Printf("Hash: %x\n", newBlock.Hash)
		//fmt.Println()
		return nil
	})
}

// 迭代器，
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}

	return bci
}

// Next returns next block starting from the tip
//反序列化并获取下一个区块的hash
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		//pow := NewProofOfWork(block)
		//fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	i.currentHash = block.PrevBlockHash

	return block
}

//数据库是否存在
func dbExists() bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}

	return true
}
// NewBlockchain creates a new Blockchain with genesis Block
//NewBlockchain() mean i can create a genesis block if the DBfile is not exist.
// i will get the data from DBfile if the DBfile is exist.
func NewBlockchain(address string) *Blockchain {


	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			fmt.Println("No existing blockchain found. Creating a new one...")
			transaction:=NewCoinbaseTX(address)
			genesis := NewGenesisBlock([]*Transaction{transaction})

			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}

			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}

			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				log.Panic(err)
			}
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	bc := Blockchain{tip, db}

	return &bc
}

//循环打印出区块链中的所有数据并进行了验证，一旦成功就说明，不管是序列化反序列化，都是成功的。
func (bc *Blockchain) printChain() {
	bci := bc.Iterator()

	for {
		block := bci.Next()
		fmt.Printf("Prev. version: %s\n", strconv.FormatInt(block.Version,10))
		fmt.Printf("Prev. hash: %x\n",block.PrevBlockHash)
		fmt.Printf("merkleroot: %s\n", block.MerkleRoot)
		fmt.Printf("time: %s\n", strconv.FormatInt(block.Timestamp,10))
		fmt.Printf("nbits: %s\n", strconv.FormatInt(block.Nbits,10))
		fmt.Printf("nonce: %s\n", strconv.FormatInt(block.Nonce,10))
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}


}


// FindUnspentTransactions returns a list of transactions containing unspent outputs
//查询未花费的交易，非常经典
func (bc *Blockchain) FindUnspentTransactions(address string) []Transaction {
	var unspentTXs []Transaction
	spentTXOs := make(map[string][]int)
	bci := bc.Iterator()

	//循环每个区块，从最后一个区块到创世区块，最后一个区块的输出一定是没有花费的。
	for {
		block := bci.Next()

		//循环区块中每个交易
		for _, tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)

		Outputs:
			//循环每个交易中的输出，如果这些输出未花费，并且能够解锁，那么既可以
			for outIdx, out := range tx.Vout {
				// Was the output spent?
				if spentTXOs[txID] != nil {
					for _, spentOut := range spentTXOs[txID] {
						if spentOut == outIdx {
							continue Outputs
						}
					}
				}

				if out.CanBeUnlockedWith(address) {
					unspentTXs = append(unspentTXs, *tx)
				}
			}

			//如果不是coinbase交易，则遍历输入，加入到已经花费的数组spentTXOs中
			if tx.IsCoinbase() == false {
				for _, in := range tx.Vin {
					if in.CanUnlockOutputWith(address) {
						inTxID := hex.EncodeToString(in.Txid)
						spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)
					}
				}
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return unspentTXs
}

// FindUTXO finds and returns all unspent transaction outputs
//返回为花费的所有输出结构体TXOutput
func (bc *Blockchain) FindUTXO(address string) []TXOutput {
	var UTXOs []TXOutput
	unspentTransactions := bc.FindUnspentTransactions(address)

	for _, tx := range unspentTransactions {
		for _, out := range tx.Vout {
			if out.CanBeUnlockedWith(address) {
				UTXOs = append(UTXOs, out)
			}
		}
	}

	return UTXOs
}

// FindSpendableOutputs finds and returns unspent outputs to reference in inputs
//返回地址address可用的金额amount的交易数组
func (bc *Blockchain) FindSpendableOutputs(address string, amount int) (int, map[string][]int) {
	unspentOutputs := make(map[string][]int)
	unspentTXs := bc.FindUnspentTransactions(address)
	accumulated := 0

Work:
	for _, tx := range unspentTXs {
		txID := hex.EncodeToString(tx.ID)

		for outIdx, out := range tx.Vout {
			if out.CanBeUnlockedWith(address) && accumulated < amount {
				accumulated += out.Value
				unspentOutputs[txID] = append(unspentOutputs[txID], outIdx)

				if accumulated >= amount {
					break Work
				}
			}
		}
	}

	return accumulated, unspentOutputs
}