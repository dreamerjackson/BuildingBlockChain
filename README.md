# BuildingBlockChain


>1、将NewBlockchain的逻辑分为了两部分，一部分是创建Blockchain，一部分是根据数据库文件获取已有的blockchain

>2、修改cli中send与getBalance的逻辑，只是获取数据库构建blockchain，不会创建新的blockchain

>3、在cli中添加CreateBlockChain的方法。

>4、



测试：
```
localhost:BuildingBlockChain jackson$ ./BuildingBlockChain createwallet
publkey = 57bd65738afa763c0185b6d60b4740e646e36518d9133537529cf46ee332856e19aba80ecf28dc7a1fc284809c55b5bfdc0b912969dfc6f4cb343102eb54c6a6
prikey  = e90f89fce0b1d1de9488fb57089c922ade930c5b52eac57f8cf326067903e60b
Your new address: 1B1oiRN6FYv7VBSwLLsWiyELYsMnK14DAb
localhost:BuildingBlockChain jackson$ ./BuildingBlockChain createwallet
publkey = d317ff4866743af6af9204da9c75e3608642e0f2d7a30ddd0e6a793aea4a700b1ff178c947f58ec56d37179bbb738538c86fe682511c2e5a0a1e900937c836cf
prikey  = ba2ca98bb251fcb8dae0f8f6afaf86e34db84020fd23c2002b9d71d4a3272c73
Your new address: 17pHuBtdV7qJPUZB47TxY9cDDpJ3vZ1zfc
localhost:BuildingBlockChain jackson$ ./BuildingBlockChain createblockchain -address 1B1oiRN6FYv7VBSwLLsWiyELYsMnK14DAb
Mining a new block

Prev. version: 2
Prev. hash:
merkleroot: abc
time: 1536735569
nbits: 111111
nonce: 13684
Hash: 0000adfd7c6b8aea151bb76e97f0733f525f4b427f8484499bb9421902f671ab
3574
PoW: true
------------------------------------------------------------

localhost:BuildingBlockChain jackson$ ./BuildingBlockChain createwallet
publkey = 57bd65738afa763c0185b6d60b4740e646e36518d9133537529cf46ee332856e19aba80ecf28dc7a1fc284809c55b5bfdc0b912969dfc6f4cb343102eb54c6a6
prikey  = e90f89fce0b1d1de9488fb57089c922ade930c5b52eac57f8cf326067903e60b
Your new address: 1B1oiRN6FYv7VBSwLLsWiyELYsMnK14DAb

localhost:BuildingBlockChain jackson$ ./BuildingBlockChain createwallet
publkey = d317ff4866743af6af9204da9c75e3608642e0f2d7a30ddd0e6a793aea4a700b1ff178c947f58ec56d37179bbb738538c86fe682511c2e5a0a1e900937c836cf
prikey  = ba2ca98bb251fcb8dae0f8f6afaf86e34db84020fd23c2002b9d71d4a3272c73
Your new address: 17pHuBtdV7qJPUZB47TxY9cDDpJ3vZ1zfc

localhost:BuildingBlockChain jackson$ ./BuildingBlockChain createblockchain -address 1B1oiRN6FYv7VBSwLLsWiyELYsMnK14DAb
Mining a new block

Prev. version: 2
Prev. hash:
merkleroot: abc
time: 1536735569
nbits: 111111
nonce: 13684
Hash: 0000adfd7c6b8aea151bb76e97f0733f525f4b427f8484499bb9421902f671ab
3574
PoW: true
------------------------------------------------------------

Done!


localhost:BuildingBlockChain jackson$ ./BuildingBlockChain createblockchain getbalance -address

localhost:BuildingBlockChain jackson$ ./BuildingBlockChain getbalance -address 1B1oiRN6FYv7VBSwLLsWiyELYsMnK14DAb
Balance of '1B1oiRN6FYv7VBSwLLsWiyELYsMnK14DAb': 100
localhost:BuildingBlockChain jackson$ ./BuildingBlockChain getbalance -address 17pHuBtdV7qJPUZB47TxY9cDDpJ3vZ1zfc
Balance of '17pHuBtdV7qJPUZB47TxY9cDDpJ3vZ1zfc': 0

localhost:BuildingBlockChain jackson$ ./BuildingBlockChain send -from 1B1oiRN6FYv7VBSwLLsWiyELYsMnK14DAb -to 17pHuBtdV7qJPUZB47TxY9cDDpJ3vZ1zfc -amount 90
Mining a new block

Success!
localhost:BuildingBlockChain jackson$ ./BuildingBlockChain getbalance -address 17pHuBtdV7qJPUZB47TxY9cDDpJ3vZ1zfc
Balance of '17pHuBtdV7qJPUZB47TxY9cDDpJ3vZ1zfc': 90
localhost:BuildingBlockChain jackson$ ./BuildingBlockChain getbalance -address 1B1oiRN6FYv7VBSwLLsWiyELYsMnK14DAb
Balance of '1B1oiRN6FYv7VBSwLLsWiyELYsMnK14DAb': 10
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