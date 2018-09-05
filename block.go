package main

import (
	"bytes"
	"strconv"
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	version       int64
	PrevBlockHash []byte
	MerkleRoot    []byte
	Timestamp     int64
	nBits         int64
	Nonce         int64
}

func (b *Block) serialize() []byte{
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	nBits := []byte(strconv.FormatInt(b.nBits, 10))
	Nonce :=[]byte(strconv.FormatInt(b.Nonce, 10))
	result := bytes.Join([][]byte{IntToHex64(b.version),b.PrevBlockHash,b.MerkleRoot,timestamp,nBits,Nonce}, []byte{})
	return result
}


func (b *Block) setHash() []byte{
	hash := sha256.Sum256(b.serialize())
	rev := sha256.Sum256(hash[:])
	return rev[:]
}

func main(){
	block := &Block{2,[]byte("abc"),[]byte("dfg"),time.Now().Unix(),111111,100}
	fmt.Printf("%x",block.serialize())
}


