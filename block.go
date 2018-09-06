package main

import (
	"bytes"
	"strconv"
	"crypto/sha256"
	"fmt"
	"encoding/hex"
)

type Block struct {
	version       int64
	PrevBlockHash []byte
	MerkleRoot    []byte
	Timestamp     int64
	nBits         int64
	Nonce         int64
	Transactions  []*Transaction
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
func	TestFuncTransactionToMerkletree(){
	txin := TXInput{[]byte{}, -1, nil}
	txout := NewTXOutput(subsidy, "first")
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{*txout}}
	tx.ID = tx.Hash()

	txin2 := TXInput{[]byte{}, -1, nil}
	txout2 := NewTXOutput(subsidy, "second")
	tx2 := Transaction{nil, []TXInput{txin2}, []TXOutput{*txout2}}
	tx2.ID = tx2.Hash()

	var Transactions  []*Transaction
	Transactions = append(Transactions, &tx,&tx2)
	fmt.Printf("%x\n",TestcreateMerkelTreeRoot(Transactions))
}
func Test_blockSerialize(){
	//block := &Block{2,[]byte("abc"),[]byte("dfg"),time.Now().Unix(),111111,100}
	//fmt.Printf("%x",block.serialize())
}
func Test_NewCoinbaseTX(){
	//newtx :=NewCoinbaseTX("tom")
	//fmt.Printf("%s",newtx.Vout)

}

func Test_NewBlockSerialize(){
	//block := &Block{2,[]byte("abc"),[]byte("dfg"),time.Now().Unix(),111111,100}
	//fmt.Printf("%x",block.serialize())

}
func main(){

	TestFuncTransactionToMerkletree()


}


