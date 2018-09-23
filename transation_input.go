package main

import "bytes"

type TXInput struct {
	Txid      []byte
	Vout      int
	Signature []byte
	Pubkey    []byte
}



// CanUnlockOutputWith checks whether the address initiated the transaction
func (in *TXInput) UsesKey(unlockingData []byte) bool {
	lockingHash := HashPubKey(in.Pubkey)

	return bytes.Compare(lockingHash, unlockingData) == 0
}

