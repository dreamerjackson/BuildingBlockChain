# BuildingBlockChain


>1、增加findUTXO(）查找并返回所有的未花费的输出。

>2、增加 type TXOutputs struct {Outputs []TXOutput} 存储交易的输出，并添加序列化和反序列化的函数。

>3、增加utxo_set文件，

>4、挖矿需要返回新区块，MineBlock(transactions []*Transaction)  *Block，在添加一个新的区块之后，应该调用updateutxo方法，更新数据库。

```
utxo_set文件

1、FindSpendableOutputs 发现可花费的输出，但是是查询的文件UTXO
2、FindUTXO(pubKeyHash []byte)     根据pubKeyHash查询，但是是查询的文件UTXO
3、Reindex()  删除原始数据库、新建数据库，调用blockchian.go中的findutxo遍历区块链从新构建utxo数据库
4、Update 更新utxo，当新挖到一个区块或者新建一个区块的时候，要跟新


```



测试：
```
localhost:BuildingBlockChain jackson$ ./BuildingBlockChain createwallet
publkey = 44bbfb3a1a81df7924a11474b8a7e5361ee3c275d82141955f4bc72bcf776fa40eef7533a786b17769c0f1fb52e9877989e4a822ebfec8103f8bbeb50959ce84
prikey  = 694edd6526bc1b27217f29c8f37ebc5de7d153511ce5798d1c18a5116909cdf1
Your new address: 1BXZ4C9X9te3eRXNYsfkJq7tyFuHpk8yep

localhost:BuildingBlockChain jackson$ ./BuildingBlockChain createwallet
publkey = 7c6cc7ac19e71157a60ff19de47e9d1e587ee98570b808bf5d1de11678b8b5b4c08d40359edeafa58e1c4eb5115b995f45b4928fdf5440aa038400c5f571e099
prikey  = f41862866a7eb34eed606fa88ca8099c45f688454d1c2cae79fa8629dfe6092f
Your new address: 13AF1uHDartNnXEQWw7j8R7mic7GJC7FwF

localhost:BuildingBlockChain jackson$ ./BuildingBlockChain createblockchain -address 1BXZ4C9X9te3eRXNYsfkJq7tyFuHpk8yep
Mining a new block

Prev. version: 2
Prev. hash:
merkleroot: abc
time: 1536808748
nbits: 111111
nonce: 32733
Hash: 000095022232e71d25c6b8f7a0bbd5630fb4c9ebbebfcba53f3d18f7152ebc60
7fdd
PoW: true
------------------------------------------------------------

Done!
localhost:BuildingBlockChain jackson$ ./BuildingBlockChain send -from 1BXZ4C9X9te3eRXNYsfkJq7tyFuHpk8yep -to 13AF1uHDartNnXEQWw7j8R7mic7GJC7FwF -amount 5
verify success
Mining a new block

Success!

localhost:BuildingBlockChain jackson$ ./BuildingBlockChain getbalance -address 1BXZ4C9X9te3eRXNYsfkJq7tyFuHpk8yep
Balance of '1BXZ4C9X9te3eRXNYsfkJq7tyFuHpk8yep': 95
localhost:BuildingBlockChain jackson$ ./BuildingBlockChain getbalance -address 13AF1uHDartNnXEQWw7j8R7mic7GJC7FwF
Balance of '13AF1uHDartNnXEQWw7j8R7mic7GJC7FwF': 5


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