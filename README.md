# BuildingBlockChain


>0、重构网络，将其就改为通道链接，所有的节点都是一个服务器，维护了所有的Client链接



```
node3000节点先建立，node3001节点链接上3000后立即发送了sendversion


		case "version":
			handleVersion(client,request, bc)//获取区块高度，如果自己的高度更高，发送sendversion，否则发送getblocks命令

		case "getblocks":
			handleGetBlocks(client,request, bc) //    发送sendInv，里面包含了所有的区块的hash头

	    case "inv":
			handleInv(client,request, bc)      //获取所有的区块的hash头，每次取一个hash，然后调用sendGetData

		case "getdata":
			handleGetData(client,request, bc)      //获取到hash的区块数据，然后进行sendBlock


		case "block":
			handleBlock(client,request, bc)   //接受到一个新的区块，addblock添加到数据库中，addblock函数中有判断，如果数据库中已经存在则不会再添加。当接收到一个后，会再次的发送下一个区块的hash，注意这种效率是很低的，可以优化。

		default:
			fmt.Println("Unknown command!")

```





>测试
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