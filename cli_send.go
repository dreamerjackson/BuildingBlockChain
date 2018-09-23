package main

import (
	"log"
	"fmt"
)

func (cli *CLI) send(from, to string, amount int, nodeID string, mineNow bool) {
	if !ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}


	bc := NewBlockchain(nodeID)
	UTXOSet := UTXOSet{bc}
	defer bc.db.Close()

	tx := NewUTXOTransaction(from, to, amount, bc,nodeID)
	//fmt.Printf("slide-----%v------------1\n",bc.VerifyTransaction(tx))
	//seriatx := DeserializeTransaction((*tx).Serialize())
	//seriatx2:= DeserializeTransaction(gobEncode(tx))
	//fmt.Printf("slide-----%v------------2\n",bc.VerifyTransaction(&seriatx))
	//fmt.Printf("slide-----%v------------3\n",bc.VerifyTransaction(&seriatx2))
	//fmt.Println(tx.String())
	//fmt.Printf("%x",gobEncode(tx))
	if mineNow {
		cbTx := NewCoinbaseTX(from, "")
		txs := []*Transaction{cbTx, tx}

		newBlock := bc.MineBlock(txs)

		UTXOSet.Update(newBlock)
	} else {
		sendTx(knownNodes[0], tx)
	}
	fmt.Println("Success!")
}

