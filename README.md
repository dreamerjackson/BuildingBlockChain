# BuildingBlockChain


>1、验证：func (bc *Blockchain) VerifyTransaction(tx *Transaction) bool 验证每一笔交易是否有效。

>2、上面调用了func (tx *Transaction) Verify(prevTXs map[string]Transaction) bool，验证交易，传递的参数是交易的输入的id，代表引用的前一笔交易->输入的交易的结构体Transaction

>3、在mineblock挖矿中，需要验证每一笔交易是否是有效的。

>4、



测试：
```
localhost:BuildingBlockChain jackson$ go build .
localhost:BuildingBlockChain jackson$ ./BuildingBlockChain send -from 1B1oiRN6FYv7VBSwLLsWiyELYsMnK14DAb -to 17pHuBtdV7qJPUZB47TxY9cDDpJ3vZ1zfc -amount 5
verify success
Mining a new block

Success!

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