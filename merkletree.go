package main

import "crypto/sha256"

//求最小值
func min(a int,b int) int{
	if(a >b){
		return b
	}else{
		return a
	}
}
//默克尔树的根节点
type MerkleTree struct {
	RootNode *MerkleNode
}

//默克尔树中的节点
type MerkleNode struct {
	Left  *MerkleNode  //左下节点
	Right *MerkleNode  //右下节点
	Data  []byte     //当前节点的哈希值
}

//生成默克尔树中的节点，如果是叶子节点，则Left，right为nil ，如果为非叶子节点，根据Left，right生成当前节点的hash
func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	mNode := MerkleNode{}

	if left == nil && right == nil {
		mNode.Data =data
	} else {
		prevHashes := append(left.Data, right.Data...)
		hh :=sha256.Sum256(prevHashes)
		hash := sha256.Sum256(hh[:])
		mNode.Data = hash[:]
	}
	mNode.Left = left
	mNode.Right = right

	return &mNode
}
//默克尔树构建
func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode

	for _, datum := range data {
		node := NewMerkleNode(nil, nil, datum)
		nodes = append(nodes, *node)
	}
	j := 0
	for nSize := len(data); nSize > 1; nSize = (nSize + 1) / 2{
		for i := 0; i < nSize; i += 2{
			i2 := min(i+1,nSize-1)
			node := NewMerkleNode(&nodes[j+i], &nodes[j+i2], nil)
			nodes = append(nodes, *node)
		}
		j += nSize
	}
	mTree := MerkleTree{&(nodes[len(nodes)-1])}
	return &mTree
}