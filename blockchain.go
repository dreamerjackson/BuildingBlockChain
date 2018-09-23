package main

import (
	"log"
	"github.com/boltdb/bolt"
	"os"

)

const dbFile = "blockchain_%s.db"
const blocksBucket = "blocks"
const genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"
// Blockchain keeps a sequence of Blocks
//区块链，保存了最近的区块的hash以及数据库对象
type Blockchain struct {
	tip []byte
	db  *bolt.DB
}



// AddBlock saves provided data as a block in the blockchain
//增加一个区块
func (bc *Blockchain) AddBlock(block *Block) {
	err := bc.db.Update(func(tx *bolt.Tx) error {
		//添加的区块在数据库中不存在
		b := tx.Bucket([]byte(blocksBucket))
		blockInDb := b.Get(block.Hash)

		if blockInDb != nil {
			return nil
		}

		//直接blockHash->blockSerialize放入数据库
		blockData := block.Serialize()
		err := b.Put(block.Hash, blockData)
		if err != nil {
			log.Panic(err)
		}

		//判断高度，如果高度更高的话，将其添加到 l ->blockHash->blockSerialize
		lastHash := b.Get([]byte("l"))
		lastBlockData := b.Get(lastHash)
		lastBlock := DeserializeBlock(lastBlockData)

		if block.Height > lastBlock.Height {
			err = b.Put([]byte("l"), block.Hash)
			if err != nil {
				log.Panic(err)
			}
			bc.tip = block.Hash
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}


//数据库是否存在
//func dbExists() bool {
//	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
//		return false
//	}
//
//	return true
//}

func dbExists(dbFile string) bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}

	return true
}








