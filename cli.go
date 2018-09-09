package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"log"
)

// CLI responsible for processing command line arguments
type CLI struct {
	bc *Blockchain
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  addblock -data BLOCK_DATA - add a block to the blockchain")
	fmt.Println("  printchain - print all the blocks of the blockchain")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 1 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock() {
	cli.bc.AddBlock()
	fmt.Println("Success!")
}

func (cli *CLI) printChain() {
	bci := cli.bc.Iterator()

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


func (cli *CLI) send(from, to string, amount int) {
	bc := NewBlockchain()
	defer bc.db.Close()
	fmt.Println("33333!")
	tx := NewUTXOTransaction(from, to, amount, bc)
	fmt.Println("444444!")
	bc.MineBlock([]*Transaction{tx})
	fmt.Println("Success!")
}


func (cli *CLI) getBalance(address string) {
	bc := NewBlockchain()
	defer bc.db.Close()

	balance := 0
	UTXOs := bc.FindUTXO(address)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}




// Run parses command line arguments and processes commands
func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	sendFrom := sendCmd.String("from", "", "Source wallet address")
	sendTo := sendCmd.String("to", "", "Destination wallet address")
	sendAmount := sendCmd.Int("amount", 0, "Amount to send")
	getBalanceAddress := getBalanceCmd.String("address", "", "The address to get balance for")
	switch os.Args[1] {
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	default:
		cli.printUsage()
		os.Exit(1)
	}

	if getBalanceCmd.Parsed() {
		if *getBalanceAddress == "" {
			getBalanceCmd.Usage()
			os.Exit(1)
		}
		cli.getBalance(*getBalanceAddress)
	}
	if addBlockCmd.Parsed() {
		cli.addBlock()
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}


	if sendCmd.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
			sendCmd.Usage()
			os.Exit(1)
		}

		cli.send(*sendFrom, *sendTo, *sendAmount)
	}
}
