# BuildingBlockChain


>0、增加网络处理server.go

>1、以前的程序，db文件的文件名都是硬编码的方式写入，现在传递了nodeID来标示不同的节点。所以在bolckchain文件中以及wallet文件中，都惊醒了大量的修改

>2、增加 GetBestHeight()、GetBlock(blockHash []byte)根据区块链的hash获取区块、GetBlockHashes() [][]byte获取区块链中存在的全部区块的hash

>3、block结构体添加高度height，创世区块高度为0，在mineblock中都需要更新

>4、修改CLI中的方法，都需要传递nodeID。

>5、addBlock在之前是没有用的，现在将cli中的addBlock删除，现在由于有了网络，接受到的区块就可以添加到区块链中。

>6、增加交易的反序列化操作DeserializeTransaction

>7、修改cli中printblockchain方法

```
// VerifyTransaction verifies transaction input signatures
func (bc *Blockchain) VerifyTransaction(tx *Transaction) bool {
添加重要一步，如果是矿工的区块，那么直接返回true
	if tx.IsCoinbase() {
		return true
	}



```



测试：
```
第一个终端：
export NODE_ID=3000
localhost:BuildingBlockChain jackson$ go build .
localhost:BuildingBlockChain jackson$ ./BuildingBlockChain createblockchain -address 18fHTVuCgUwz1MHCVCb69jAark5KGg5QEPMining a new block

注意在这个地方，将生成的blockchain_3000.db数据库复制一个副本到blockchain_3001.db


Prev. version: 2
Prev. hash:
merkleroot: abc
time: 1536984640
nbits: 111111
nonce: 45673
Hash: 0000c15394a826f86b4d538e48bc2a053882dd9e9caaffedf35c29c0c8be06cb
b269
PoW: true
------------------------------------------------------------

Done!

localhost:BuildingBlockChain jackson$ ./BuildingBlockChain getbalance -address 18fHTVuCgUwz1MHCVCb69jAark5KGg5QEPBalance of '18fHTVuCgUwz1MHCVCb69jAark5KGg5QEP': 100
localhost:BuildingBlockChain jackson$ ./BuildingBlockChain send -from 18fHTVuCgUwz1MHCVCb69jAark5KGg5QEP  -to 1MKQN5oVYwwWofQe3n45V6mupRsWQ1QmD5 -amount 10 -mine
verify success
verify success
Mining a new block

Success!
localhost:BuildingBlockChain jackson$ ./BuildingBlockChain getbalance -address 18fHTVuCgUwz1MHCVCb69jAark5KGg5QEPBalance of '18fHTVuCgUwz1MHCVCb69jAark5KGg5QEP': 190
localhost:BuildingBlockChain jackson$ ./BuildingBlockChain startnode
Starting node 3000
Received version command
Received getblocks command
Received getdata command
Received getdata command

```


```
第二个终端：

export NODE_ID=3001

没有同步前，数据库中存储的金额是100
bogon:BuildingBlockChain jackson$ ./BuildingBlockChain getbalance -address 18fHTVuCgUwz1MHCVCb69jAark5KGg5QEP
Balance of '18fHTVuCgUwz1MHCVCb69jAark5KGg5QEP': 100



同步后，接受到了数据，并将其添加到了数据库中。
bogon:BuildingBlockChain jackson$ ./BuildingBlockChain startnode
Starting node 3001
Received version command
Received inv command
Recevied inventory with 2 block
Received block command
Recevied a new block!
Added block 0000504829f3f877acf18ede4d68193cf2be3c1476c30095550a1b1015181ad8
Received block command
Recevied a new block!
Added block 0000c15394a826f86b4d538e48bc2a053882dd9e9caaffedf35c29c0c8be06cb
^C
bogon:BuildingBlockChain jackson$ ./BuildingBlockChain getbalance -address 18fHTVuCgUwz1MHCVCb69jAark5KGg5QEP
Balance of '18fHTVuCgUwz1MHCVCb69jAark5KGg5QEP': 190


```




>测试2：
```
#3000

bogon:BuildingBlockChain jackson$  export NODE_ID=3000
bogon:BuildingBlockChain jackson$ ./BuildingBlockChain createblockchain -address 18fHTVuCgUwz1MHCVCb69jAark5KGg5QEP
bogon:BuildingBlockChain jackson$ ./BuildingBlockChain startnode

#3002
bogon:BuildingBlockChain jackson$ ./BuildingBlockChain startnode -miner 18fHTVuCgUwz1MHCVCb69jAark5KGg5QEP


#3001

bogon:BuildingBlockChain jackson$  export NODE_ID=3001
bogon:BuildingBlockChain jackson$ ./BuildingBlockChain send -from 18fHTVuCgUwz1MHCVCb69jAark5KGg5QEP  -to 1MKQN5oVYwwWofQe3n45V6mupRsWQ1QmD5 -amount 1
bogon:BuildingBlockChain jackson$ ./BuildingBlockChain getbalance -address 18fHTVuCgUwz1MHCVCb69jAark5KGg5QEP
bogon:BuildingBlockChain jackson$ ./BuildingBlockChain printchain
```

```
交 流 群 名 称：
Go底层公链
交 流 群 号：
713385260
```

```
情深不寿
强极则辱
谦谦君子
温润如玉
```