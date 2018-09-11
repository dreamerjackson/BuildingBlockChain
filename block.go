package main

import (
	"bytes"
	"strconv"
	"crypto/sha256"
	"fmt"
	"encoding/hex"
	"time"
	"encoding/gob"
	"log"
)
//序列化时，使用了encoding/gob，切记必须要大些
type Block struct {
	Version       int64
	PrevBlockHash []byte
	MerkleRoot    []byte
	Timestamp     int64
	Nbits         int64
	Nonce         int64
	Transactions  []*Transaction
	Hash          []byte
}


//废弃
func (b *Block) serialize() []byte{
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	nBits := []byte(strconv.FormatInt(b.Nbits, 10))
	Nonce :=[]byte(strconv.FormatInt(b.Nonce, 10))
	result := bytes.Join([][]byte{IntToHex64(b.Version),b.PrevBlockHash,b.MerkleRoot,timestamp,nBits,Nonce}, []byte{})
	return result
}


// Serialize serializes the block
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// DeserializeBlock deserializes a block
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}

func (b *Block) setHash() []byte{
	hash := sha256.Sum256(b.serialize())
	rev := sha256.Sum256(hash[:])
	return rev[:]
}

// create MerkleRoot though transactions
func (b * Block ) createMerkelTreeRoot(Transactions  []*Transaction){
	var transactions [][]byte

	for _, tx := range Transactions {
		transactions = append(transactions, tx.Serialize())
	}
	mTree := NewMerkleTree(transactions)

	b.MerkleRoot= mTree.RootNode.Data
}


func testMerkleTree(){

	//https://www.blockchain.com/btc/block/00000000000090ff2791fe41d80509af6ffbd6c5b10294e29cdf1b603acab92c

	data1,_:=hex.DecodeString("6b6a4236fb06fead0f1bd7fc4f4de123796eb51675fb55dc18c33fe12e33169d")
	data2,_:=hex.DecodeString("2af6b6f6bc6e613049637e32b1809dd767c72f912fef2b978992c6408483d77e")
	data3,_:=hex.DecodeString("6d76d15213c11fcbf4cc7e880f34c35dae43f8081ef30c6901f513ce41374583")
	data4,_:=hex.DecodeString("08c3b50053b010542dca85594af182f8fcf2f0d2bfe8a806e9494e4792222ad2")
	data5,_:=hex.DecodeString("612d035670b7b9dad50f987dfa000a5324ecb3e08745cfefa10a4cefc5544553")
	data6:= reverse2(data1)
	data7:= reverse2(data2)
	data8:= reverse2(data3)
	data9:= reverse2(data4)
	data10:= reverse2(data5)

	hehe := [][]byte{
		data6,
		data7,
		data8,
		data9,
		data10,
	}

	result := (*NewMerkleTree(hehe).RootNode).Data
	rev := reverse2(result)
	fmt.Printf("result=%x\n", rev)
}

// create MerkleRoot though transactions
func  TestcreateMerkelTreeRoot(Transactions  []*Transaction) []byte{
	var transactions [][]byte

	for _, tx := range Transactions {
		transactions = append(transactions, tx.Serialize())
	}
	mTree := NewMerkleTree(transactions)
	return mTree.RootNode.Data
	//b.MerkleRoot= mTree.RootNode.Data
}

//产生初始区块,传入了第一笔coinbase交易
func NewGenesisBlock(transactions []*Transaction) *Block {
	block :=&Block{int64(2),[]byte{},[]byte("abc"),time.Now().Unix(),111111,100,transactions,[]byte{}}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	//fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	fmt.Printf("Prev. version: %s\n", strconv.FormatInt(block.Version,10))
	fmt.Printf("Prev. hash: %x\n",block.PrevBlockHash)
	fmt.Printf("merkleroot: %s\n", block.MerkleRoot)
	fmt.Printf("time: %s\n", strconv.FormatInt(block.Timestamp,10))
	fmt.Printf("nbits: %s\n", strconv.FormatInt(block.Nbits,10))
	fmt.Printf("nonce: %s\n", strconv.FormatInt(block.Nonce,10))
	fmt.Printf("Hash: %x\n", block.Hash)
	fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	fmt.Printf("------------------------------------------------------------\n")
	fmt.Println()
	return block
	}

func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block{
	block :=&Block{2,prevBlockHash,[]byte("dfg"),time.Now().Unix(),111111,0,transactions,[]byte{}}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	return block

}


func TestSerialize(){
	//k :=&Blocktest{[]byte("jhg"),[]byte("abc"),time.Now().Unix(),100,[]*Transaction{},[]byte{}}
	//block := DeserializeBlock2(k.Serialize())
	////fmt.Printf("Prev. version: %s\n", strconv.FormatInt(block.version,10))
	//fmt.Printf("Prev. hash: %x\n",block.PrevBlockHash)
	//fmt.Printf("merkleroot: %s\n", block.MerkleRoot)
	//fmt.Printf("time: %s\n", strconv.FormatInt(block.Timestamp,10))
	////fmt.Printf("nbits: %s\n", strconv.FormatInt(block.nBits,10))
	//fmt.Printf("nonce: %s\n", strconv.FormatInt(block.Version,10))
	//fmt.Printf("Hash: %x\n", block.Hash)
	////fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	//fmt.Printf("------------------------------------------------------------\n")
	//fmt.Println()
}

func TestBoltDB(){
	//NewBlockchain() mean i can create a genesis block if the DBfile is not exist.
	// i will get the data from DBfile if the DBfile is exist.
	//blockchian :=NewBlockchain()  已经修改
	//blockchian.AddBlock()
	//blockchian.AddBlock()
	//blockchian.printChain()

	}
