package main
//
//import (
//	"crypto/sha256"
//	"encoding/hex"
//	"fmt"
//	"bytes"
//)
////求最小值
//func min(a int,b int) int{
//	if(a >b){
//		return b
//	}else{
//		return a
//	}
//}
////默克尔树的根节点
//type MerkleTree struct {
//	RootNode *MerkleNode
//}
//
////默克尔树中的节点
//type MerkleNode struct {
//	Left  *MerkleNode  //左下节点
//	Right *MerkleNode  //右下节点
//	Data  []byte     //当前节点的哈希值
//}
//
////生成默克尔树中的节点，如果是叶子节点，则Left，right为nil ，如果为非叶子节点，根据Left，right生成当前节点的hash
//func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
//	mNode := MerkleNode{}
//
//	if left == nil && right == nil {
//		mNode.Data =data
//	} else {
//		prevHashes := append(left.Data, right.Data...)
//		hh :=sha256.Sum256(prevHashes)
//		hash := sha256.Sum256(hh[:])
//		mNode.Data = hash[:]
//	}
//	mNode.Left = left
//	mNode.Right = right
//
//	return &mNode
//}
////默克尔树构建
//func NewMerkleTree2(data [][]byte) *MerkleTree {
//	var nodes []MerkleNode
//
//	for _, datum := range data {
//		node := NewMerkleNode(nil, nil, datum)
//		nodes = append(nodes, *node)
//	}
//	j := 0
//	for nSize := len(data); nSize > 1; nSize = (nSize + 1) / 2{
//		for i := 0; i < nSize; i += 2{
//			i2 := min(i+1,nSize-1)
//			node := NewMerkleNode(&nodes[j+i], &nodes[j+i2], nil)
//			nodes = append(nodes, *node)
//		}
//		j += nSize
//	}
//	mTree := MerkleTree{&(nodes[len(nodes)-1])}
//	return &mTree
//}
//
////大小端转换
//func reverse(data []byte) []byte{
//	var s [][]byte
//	for i:=len(data);i>0;i--{
//		data1 :=data[i-1:i]
//		s = append(s,data1)
//	}
//	sep := []byte("")
//	result :=bytes.Join(s,sep)
//	return result
//}
//
//func testMerkleTree(){
//
//	//https://www.blockchain.com/btc/block/00000000000090ff2791fe41d80509af6ffbd6c5b10294e29cdf1b603acab92c
//
//	data1,_:=hex.DecodeString("6b6a4236fb06fead0f1bd7fc4f4de123796eb51675fb55dc18c33fe12e33169d")
//	data2,_:=hex.DecodeString("2af6b6f6bc6e613049637e32b1809dd767c72f912fef2b978992c6408483d77e")
//	data3,_:=hex.DecodeString("6d76d15213c11fcbf4cc7e880f34c35dae43f8081ef30c6901f513ce41374583")
//	data4,_:=hex.DecodeString("08c3b50053b010542dca85594af182f8fcf2f0d2bfe8a806e9494e4792222ad2")
//	data5,_:=hex.DecodeString("612d035670b7b9dad50f987dfa000a5324ecb3e08745cfefa10a4cefc5544553")
//	data6:= reverse(data1)
//	data7:= reverse(data2)
//	data8:= reverse(data3)
//	data9:= reverse(data4)
//	data10:= reverse(data5)
//
//	hehe := [][]byte{
//		data6,
//		data7,
//		data8,
//		data9,
//		data10,
//	}
//
//	result := (*NewMerkleTree2(hehe).RootNode).Data
//	rev := reverse(result)
//	fmt.Printf("result=%x\n", rev)
//}
//
//
//func main(){
//	testMerkleTree();
//}