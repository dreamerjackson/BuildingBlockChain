# BuildingBlockChain


>1、添加了数字签名，首先修改了input的结构体，增加了publickey，
修改了NewUTXOTransaction(from, to string, amount int, bc *Blockchain)  1、根据from地址获取到wallet对象，构建input。根据to地址拿到output中的pubkeyhash，构建output。

>2、NewCoinbaseTX(to, data string),传递一个const genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

>3、修改FindUTXO(pubKeyHash []byte)众多函数，参数改为pubkeyhash。

>4、修改CanBeUnlockedWith(pubKeyHash []byte)，能够解锁的条件是pubkeyhash相同



测试：
、、、
localhost:BuildingBlockChain jackson$ ./BuildingBlockChain getbalance -address 1LrW1JzBooC5uELS4My4t1Kv1K2xgFnZpY
Balance of '1LrW1JzBooC5uELS4My4t1Kv1K2xgFnZpY': 0
localhost:BuildingBlockChain jackson$ ./BuildingBlockChain getbalance -address 1Fa8U2ETHCQQHqhLwHS91xuURo8DXakRY1
Balance of '1Fa8U2ETHCQQHqhLwHS91xuURo8DXakRY1': 6
localhost:BuildingBlockChain jackson$ go build .
localhost:BuildingBlockChain jackson$ ./BuildingBlockChain createwallet
publkey = 5c1dd8d5901ae649663f79ab0a249c9fd8c065a1f2b1f85223d1eda96e46d402dfae4f2352867694be35b5c314f74fe79215999b25ab14d549aa7c04e5ce08da
prikey  = 2140b919fb8f49c5ee3d63634a75150bdd557a3496f7a8844eb8f7cb38e70ef3
Your new address: 1P81XQuUWnJVZDWai8yXoxixTNvfkP3p3o
localhost:BuildingBlockChain jackson$ ./BuildingBlockChain createwallet
publkey = 70521e696d1b00baba808e9c3cd0d7157462d9a61184c947c5a4ceedc71abb55650377c6ed005639ccf6d5de2081aacb3c73ec961ba708b9098499a0c3c8adda
prikey  = 0e3870a63650a8422446e5326627232958811b915ca66cfb2022e87f9810d8fe
Your new address: 14Si1pVjDur6mH7X5ssoFE43Cn41qtgmK1
localhost:BuildingBlockChain jackson$ ./BuildingBlockChain send -from 1P81XQuUWnJVZDWai8yXoxixTNvfkP3p3o  -to 14Si1pVjDur6mH7X5ssoFE43Cn41qtgmK1  -amount 6
No existing blockchain found. Creating a new one...
Mining a new block

Prev. version: 2
Prev. hash:
merkleroot: abc
time: 1536649027
nbits: 111111
nonce: 116577
Hash: 0000f2a38e4923c50daf176793293cbba06f3af841d2ab916a09be20bbad16de
1c761
PoW: true
------------------------------------------------------------

Mining a new block

Success!
localhost:BuildingBlockChain jackson$ ./BuildingBlockChain getbalance -address 1Fa8U2ETHCQQHqhLwHS91xuURo8DXakRY1
Balance of '1Fa8U2ETHCQQHqhLwHS91xuURo8DXakRY1': 0
localhost:BuildingBlockChain jackson$ ./BuildingBlockChain getbalance -address 1P81XQuUWnJVZDWai8yXoxixTNvfkP3p3o
Balance of '1P81XQuUWnJVZDWai8yXoxixTNvfkP3p3o': 94
、、、

1. 创建blockchain，创建创世区块、创世交易的的功能放在了cli.send中，未来要分离
2. ./BuildingBlockChain createwallet创建两个钱包
3. ./BuildingBlockChain send -from 1P81XQuUWnJVZDWai8yXoxixTNvfkP3p3o  -to 14Si1pVjDur6mH7X5ssoFE43Cn41qtgmK1  -amount 6转账了100元给第一个地址，并执行了转账6元。
4. 转账就涉及到utxo的查询，构建交易的时候就涉及到对交易的数字签名。

、、、
交 流 群 名 称：
Go底层公链
交 流 群 号：
713385260
、、、
