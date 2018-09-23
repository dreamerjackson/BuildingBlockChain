package main

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
	"strconv"
)

// BlockchainIterator is used to iterate over blockchain blocks
//迭代器，用于循环区块
type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
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
