# Event log

Event 是在智能合约执行期间可能触发的事件，特别是在token合约中，比如token 被 transfer的时候。  

Event log 是被存储在 transaction receipts 里。

我们可以从 blockchain 中读取 Event, 也可以订阅 event。

## 订阅日志事件

## 读取事件日志

## 读取ERC20 Token事件日志

## 读取 0x 协议事件日志

定义智能合约 `exchange.sol`

生成go文件

```bash
solc --abi Exchange.sol -o build
abigen --abi="build/Exchange.abi" --pkg=event --out=exchange.go
```