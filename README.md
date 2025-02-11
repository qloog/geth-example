# go-ethereum example

该项目是 `go-ethereum` 的学习笔记，包括 `go-ethereum` 的基本用法，如何和链上进行交互，包括交易，智能合约，事件日志，签名等

## 下载

```bash
# path to your dir
git clone github.com/qloog/geth-example
```

## 运行
```bash
cd geth-example

go mod tidy
make run
```

## 主体内容

## 客户端

- Ganache

### 账户

### 交易

### 智能合约

### 事件日志

- 订阅事件日志
- 读取事件日志
- 读取ERC20 Token事件日志
- 读取 0x 协议事件日志

### 签名

- 生成签名
- 验证签名

### 测试

本地模拟client进行测试，比如：进行交易测试

### Swarm

- 配置 Swarm
- 上传文件到 Swarm
- 从 Swarm 下载文件

## References

- Ethereum Development with Go: https://goethereumbook.org/en/
- 面向 Go 开发者的以太坊: https://ethereum.org/zh/developers/docs/programming-languages/golang/
- 以太坊技术与实现: https://learnblockchain.cn/books/geth/
- https://github.com/ethereum/go-ethereum
- http://truffleframework.com/ganache/
- Ethereum HD Wallet derivations in Go: https://github.com/miguelmota/go-ethereum-hdwallet
- https://github.com/Luboy23/foundry_advanced_turtorial
- WTF 学院: https://www.wtf.academy/zh 
