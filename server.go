package main

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"time"
)

type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

type Client struct {
	socket net.Conn
	data   chan []byte
}

func (manager *ClientManager) start() {
	for {
		select {
		case connection := <-manager.register:
			manager.clients[connection] = true
			fmt.Println("Added new connection!")
		case connection := <-manager.unregister:
			if _, ok := manager.clients[connection]; ok {
				close(connection.data)
				delete(manager.clients, connection)
				fmt.Println("A connection has terminated!")
			}
		case message := <-manager.broadcast:
			for connection := range manager.clients {
				select {
				case connection.data <- message:
				default:
					close(connection.data)
					delete(manager.clients, connection)
				}
			}
		}
	}
}


func (manager *ClientManager) send(client *Client) {
	defer client.socket.Close()
	for {
		select {
		case message, ok := <-client.data:
			if !ok {
				return
			}
			client.socket.Write(message)
		}
	}
}

func (manager *ClientManager) receive(client *Client,bc*Blockchain) {
	for {
		fmt.Println("准备接收")
		//request, err := ioutil.ReadAll(client.socket)
		request := make([]byte, 4096)
		_, err := client.socket.Read(request)
		fmt.Println("recive 完毕")
		if err != nil {
			manager.unregister <- client
			client.socket.Close()
			log.Panic(err)
			break
		}
		fmt.Println("recive 完毕2")
		command := bytesToCommand(request[:commandLength])
		fmt.Printf("Received %s command\n", command)

		switch command {
		case "version":
			handleVersion(client,request, bc)

		case "block":
			handleBlock(client,request, bc)
		case "inv":
			handleInv(client,request, bc)
		case "getblocks":
			handleGetBlocks(client,request, bc)
		case "getdata":
			handleGetData(client,request, bc)
		default:
			fmt.Println("Unknown command!")
		}


		//message := make([]byte, 4096)
		//length, err := client.socket.Read(message)
		//if err != nil {
		//	manager.unregister <- client
		//	client.socket.Close()
		//	break
		//}
		//if length > 0 {
		//	fmt.Println("RECEIVED: " + string(message))
		//	manager.broadcast <- message
		//}
	}
}



func StartServer(nodeID, minerAddress string) {

	nodeAddress = fmt.Sprintf("localhost:%s", nodeID)
	miningAddress = minerAddress


	fmt.Println("Starting server...")
	listener, error := net.Listen("tcp", nodeAddress)
	if error != nil {
		fmt.Println(error)
	}

	bc := NewBlockchain(nodeID)

	manager := ClientManager{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
	go manager.start()



	if nodeAddress != knownNodes[0] {
		connection, error := net.Dial("tcp", knownNodes[0])
		client := &Client{socket: connection, data: make(chan []byte)}
		manager.register <- client
		if error != nil {
			fmt.Println(error)
		}
		time.Sleep(time.Second*5)
		fmt.Println("1111111111111111111111111111111")
		sendVersion(client,knownNodes[0],bc)
		fmt.Println("2222222222222222222222222222")
		go manager.receive(client,bc)
		go manager.send(client)
	}







	for {
		connection, _ := listener.Accept()
		if error != nil {
			fmt.Println(error)
		}
		client := &Client{socket: connection, data: make(chan []byte)}
		manager.register <- client
		go manager.receive(client,bc)
		go manager.send(client)
	}
}



























const protocol = "tcp"
const nodeVersion = 1
const commandLength = 12

var nodeAddress string
var miningAddress string
var knownNodes = []string{"localhost:3000"}
var blocksInTransit = [][]byte{}
var mempool = make(map[string]Transaction)




































type addr struct {
	AddrList []string
}

type block struct {
	AddrFrom string
	Block    []byte
}

type getblocks struct {
	AddrFrom string
}

type getdata struct {
	AddrFrom string
	Type     string
	ID       []byte
}

type inv struct {
	AddrFrom string
	Type     string
	Items    [][]byte
}

type tx struct {
	AddFrom     string
	Transaction []byte
}

type verzion struct {
	Version    int
	BestHeight int
	AddrFrom   string
}

func commandToBytes(command string) []byte {
	var bytes [commandLength]byte

	for i, c := range command {
		bytes[i] = byte(c)
	}

	return bytes[:]
}

func bytesToCommand(bytes []byte) string {
	var command []byte

	for _, b := range bytes {
		if b != 0x0 {
			command = append(command, b)
		}
	}

	return fmt.Sprintf("%s", command)
}

func extractCommand(request []byte) []byte {
	return request[:commandLength]
}

//func requestBlocks() {
//	for _, node := range knownNodes {
//		sendGetBlocks(node)
//	}
//}

//func sendAddr(client *Client,address string) {
//	nodes := addr{knownNodes}
//	nodes.AddrList = append(nodes.AddrList, nodeAddress)
//	payload := gobEncode(nodes)
//	request := append(commandToBytes("addr"), payload...)
//
//	sendData(client,address, request)
//}

func sendBlock(client *Client,addr string, b *Block) {
	data := block{nodeAddress, b.Serialize()}
	payload := gobEncode(data)
	request := append(commandToBytes("block"), payload...)

	sendData(client,addr, request)
}

func sendData(client *Client,addr string, data []byte) {

	_, err:=client.socket.Write(data)

	if err != nil {
		log.Panic(err)
	}
}

func sendInv(client *Client,address, kind string, items [][]byte) {
	inventory := inv{nodeAddress, kind, items}
	payload := gobEncode(inventory)
	request := append(commandToBytes("inv"), payload...)

	sendData(client,address, request)
}






func sendGetBlocks(client *Client,address string) {
	payload := gobEncode(getblocks{nodeAddress})
	request := append(commandToBytes("getblocks"), payload...)

	sendData(client,address, request)
}

func sendGetData(client *Client,address, kind string, id []byte) {
	payload := gobEncode(getdata{nodeAddress, kind, id})
	request := append(commandToBytes("getdata"), payload...)

	sendData(client,address, request)
}

func sendTx(client *Client,addr string, tnx *Transaction) {
	data := tx{nodeAddress, tnx.Serialize()}
	payload := gobEncode(data)
	request := append(commandToBytes("tx"), payload...)

	sendData(client,addr, request)
}

func sendVersion(client *Client,addr string, bc *Blockchain) {
	bestHeight := bc.GetBestHeight()
	payload := gobEncode(verzion{nodeVersion, bestHeight, nodeAddress})

	request := append(commandToBytes("version"), payload...)

	sendData(client,addr, request)
	fmt.Println("发送完毕")
}

//func handleAddr(request []byte) {
//	var buff bytes.Buffer
//	var payload addr
//
//	buff.Write(request[commandLength:])
//	dec := gob.NewDecoder(&buff)
//	err := dec.Decode(&payload)
//	if err != nil {
//		log.Panic(err)
//	}
//
//	knownNodes = append(knownNodes, payload.AddrList...)
//	fmt.Printf("There are %d known nodes now!\n", len(knownNodes))
//	requestBlocks()
//}

func handleBlock(client *Client,request []byte, bc *Blockchain) {
	var buff bytes.Buffer
	var payload block

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	blockData := payload.Block
	block := DeserializeBlock(blockData)

	fmt.Println("Recevied a new block!")
	bc.AddBlock(block)

	fmt.Printf("Added block %x\n", block.Hash)

	if len(blocksInTransit) > 0 {
		blockHash := blocksInTransit[0]
		sendGetData(client,payload.AddrFrom, "block", blockHash)

		blocksInTransit = blocksInTransit[1:]
	} else {
		UTXOSet := UTXOSet{bc}
		UTXOSet.Reindex()
	}
}

func handleInv(client *Client,request []byte, bc *Blockchain) {
	var buff bytes.Buffer
	var payload inv

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Recevied inventory with %d %s\n", len(payload.Items), payload.Type)

	if payload.Type == "block" {
		blocksInTransit = payload.Items

		blockHash := payload.Items[0]
		sendGetData(client,payload.AddrFrom, "block", blockHash)

		newInTransit := [][]byte{}
		for _, b := range blocksInTransit {
			if bytes.Compare(b, blockHash) != 0 {
				newInTransit = append(newInTransit, b)
			}
		}
		blocksInTransit = newInTransit
	}

	if payload.Type == "tx" {
		txID := payload.Items[0]

		if mempool[hex.EncodeToString(txID)].ID == nil {
			sendGetData(client,payload.AddrFrom, "tx", txID)
		}
	}
}

func handleGetBlocks(client *Client,request []byte, bc *Blockchain) {
	var buff bytes.Buffer
	var payload getblocks

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	blocks := bc.GetBlockHashes()
	sendInv(client,payload.AddrFrom, "block", blocks)
}

func handleGetData(client *Client,request []byte, bc *Blockchain) {
	var buff bytes.Buffer
	var payload getdata

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	if payload.Type == "block" {
		block, err := bc.GetBlock([]byte(payload.ID))
		if err != nil {
			return
		}

		sendBlock(client,payload.AddrFrom, &block)
	}

	if payload.Type == "tx" {
		txID := hex.EncodeToString(payload.ID)
		tx := mempool[txID]
//fmt.Printf("\nSignature================%x\n",tx.Vin[0].Signature)
		sendTx(client,payload.AddrFrom, &tx)
		// delete(mempool, txID)
	}
}

func handleTx(client *Client,request []byte, bc *Blockchain) {
	var buff bytes.Buffer
	var payload tx

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	txData := payload.Transaction
	tx := DeserializeTransaction(txData)
	fmt.Println(tx.String())
	fmt.Printf("\n%x  ||||||fu1\n",txData)
	mempool[hex.EncodeToString(tx.ID)] = tx



	if nodeAddress == knownNodes[0] {

		var txs []*Transaction
		for id := range mempool {
			tx := mempool[id]
			if bc.VerifyTransaction(&tx) {
				txs = append(txs, &tx)
			}
		}
		for _, node := range knownNodes {
			if node != nodeAddress && node != payload.AddFrom {
				sendInv(client,node, "tx", [][]byte{tx.ID})
			}
		}
	} else {

		if len(mempool) >= 1 && len(miningAddress) > 0 {

		MineTransactions:
			var txs []*Transaction

			for id := range mempool {
				tx := mempool[id]
				if bc.VerifyTransaction(&tx) {
					txs = append(txs, &tx)
				}
			}

			if len(txs) == 0 {
				fmt.Println("All transactions are invalid! Waiting for new ones...")
				return
			}

			cbTx := NewCoinbaseTX(miningAddress, "")
			txs = append(txs, cbTx)

			newBlock := bc.MineBlock(txs)

			UTXOSet := UTXOSet{bc}
			UTXOSet.Reindex()

			fmt.Println("New block is mined!")

			for _, tx := range txs {
				txID := hex.EncodeToString(tx.ID)
				delete(mempool, txID)
			}

			for _, node := range knownNodes {
				if node != nodeAddress {
					sendInv(client,node, "block", [][]byte{newBlock.Hash})
				}
			}

			if len(mempool) > 0 {
				goto MineTransactions
			}
		}
	}
}

func handleVersion(client *Client,request []byte, bc *Blockchain) {
	var buff bytes.Buffer
	var payload verzion

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	myBestHeight := bc.GetBestHeight()
	foreignerBestHeight := payload.BestHeight

	if myBestHeight < foreignerBestHeight {
		sendGetBlocks(client,payload.AddrFrom)
	} else if myBestHeight > foreignerBestHeight {
		sendVersion(client,payload.AddrFrom, bc)
	}

	// sendAddr(payload.AddrFrom)
	if !nodeIsKnown(payload.AddrFrom) {
		knownNodes = append(knownNodes, payload.AddrFrom)
	}
}

//func handleConnection(conn net.Conn, bc *Blockchain) {
//	request, err := ioutil.ReadAll(conn)
//	if err != nil {
//		log.Panic(err)
//	}
//	command := bytesToCommand(request[:commandLength])
//	fmt.Printf("Received %s command\n", command)
//
//	switch command {
//	case "addr":
//		handleAddr(request)
//	case "block":
//		handleBlock(request, bc)
//	case "inv":
//		handleInv(request, bc)
//	case "getblocks":
//		handleGetBlocks(request, bc)
//	case "getdata":
//		handleGetData(request, bc)
//	case "tx":
//		handleTx(request, bc)
//	case "version":
//		handleVersion(request, bc)
//	default:
//		fmt.Println("Unknown command!")
//	}
//
//	conn.Close()
//}

// StartServer starts a node
//func StartServer(nodeID, minerAddress string) {
//	nodeAddress = fmt.Sprintf("localhost:%s", nodeID)
//	miningAddress = minerAddress
//	ln, err := net.Listen(protocol, nodeAddress)
//	if err != nil {
//		log.Panic(err)
//	}
//	defer ln.Close()
//
//	bc := NewBlockchain(nodeID)
//
//	if nodeAddress != knownNodes[0] {
//		sendVersion(knownNodes[0], bc)
//	}
//
//	for {
//		conn, err := ln.Accept()
//		if err != nil {
//			log.Panic(err)
//		}
//		go handleConnection(conn, bc)
//	}
//}

func gobEncode(data interface{}) []byte {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(data)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func nodeIsKnown(addr string) bool {
	for _, node := range knownNodes {
		if node == addr {
			return true
		}
	}

	return false
}
