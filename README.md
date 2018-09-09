# BuildingBlockChain


>1、产生初始区块,传入了第一笔coinbase交易func NewGenesisBlock(transactions []*Transaction) *Block，第一笔交易的矿工是tom

>2、blockchain.go中加入了查询UTXO的操作，循环遍历

>3、cli中加入了查询金额和转账的命令

>4、根据提供的交易，开始挖矿MineBlock（）

>5、根据发送者，接受者，金额，创建一笔新的交易NewUTXOTransaction（）

测试：
1、新建区块链NewBlockchain（），如果没有文件，就会初始化，构建创世区块NewGenesisBlock，创世区块的第一笔交易为给tom的100元
2、./BuildingBlockChain getbalance -address tom 获取到tom100元
3、./BuildingBlockChain send -from tom -to Helen -amount 2。
    首先判断tom确实有100元，接下来NewUTXOTransaction函数创建一笔交易，MineBlock（）将这笔交易挖矿加入到区块中。

    $ ./BuildingBlockChain getbalance -address tom
    Balance of 'tom': 98
    $ ./BuildingBlockChain getbalance -address Helen
    Balance of 'Helen': 2

查看对比：
https://github.com/dreamerjackson/BuildingBlockChain/compare