package main

import (
	"math/big"
	"math"
	"bytes"
	"fmt"
	"crypto/sha256"
)

const maxNonce = math.MaxInt64

const targetBits = 16

// ProofOfWork represents a proof-of-work
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork builds and returns a ProofOfWork
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

//serialize the block
func (pow *ProofOfWork) prepareData(nonce int64) []byte {
	data := bytes.Join(
		[][]byte{
			IntToHex64(pow.block.Version),
			pow.block.PrevBlockHash,
			pow.block.MerkleRoot,
			IntToHex64(pow.block.Timestamp),
			IntToHex64(pow.block.Nbits),
			IntToHex64(nonce),
		},
		[]byte{},
	)

	return data
}

// Run performs a proof-of-work
func (pow *ProofOfWork) Run() (int64, []byte) {
	var hashInt big.Int
	var hash [32]byte

	var nonce int64;
	nonce = 0

	fmt.Printf("Mining a new block")

	//continue the nonce until the hash result <= targetBits
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		firsthash := sha256.Sum256(data)
		hash = sha256.Sum256(firsthash[:])
		//if nonce==5{
		//	fmt.Printf("\rnonce:%x,%x\n", nonce,hash)
		//}

		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

// Validate validates block's PoW
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	firsthash := sha256.Sum256(data)
	hash := sha256.Sum256(firsthash[:])

	fmt.Printf("\r%x\n", pow.block.Nonce)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
