# Swarm

Swarm 是以太坊的去中心化分布式存储解决方案，类似IPFS。
Swarm 是一个点对点数据共享网络，其中文件通过其内容的哈希值来寻址。

## 配置 Swarm

为了运行 Swarm, 需要安装 `geth` 和 `bzzd`

> bzzd 是 Swarm 守护进程

```bash
go install github.com/ethereum/go-ethereum/cmd/geth@latest
go install github.com/ethersphere/swarm/cmd/swarm@latest
```

使用 `geth` 生成一个 account

```bash
➜  geth-example git:(main) ✗ geth account new
INFO [12-19|19:08:10.010] Maximum peer count                       ETH=50 total=50
Your new account is locked with a password. Please give a password. Do not forget this password.
Password: 
Repeat password: 

Your new key was generated

Public address of the key:   0xD090303b2494385D9451DB9e2D9BE080740fc1f8
Path of the secret key file: /Users/xxxxx

- You can share your public address with anyone. Others need it to interact with you.
- You must NEVER share the secret key with anyone! The key controls access to your funds!
- You must BACKUP your key file! Without the key, it's impossible to access account funds!
- You must REMEMBER your password! Without the password, it's impossible to decrypt the key!
```

将刚刚生成的公钥导出到环境变量 

```bash
export BZZKEY=0xD090303b2494385D9451DB9e2D9BE080740fc1f8
```

运行 Swarm

```bash
swarm --bzzaccount $BZZKEY
➜  go-eth git:(main) ✗ swarm --bzzaccount $BZZKEY
INFO [12-19|19:14:12.076] Maximum peer count                       ETH=50 LES=0 total=50
Unlocking swarm account 0xD090303b2494385D9451DB9e2D9BE080740fc1f8 [1/3]
Passphrase: 
INFO [12-19|19:14:28.637] Starting peer-to-peer node               instance=swarm/v0.5.8/darwin-arm64/go1.22.8
INFO [12-19|19:14:28.706] New local node record                    seq=1 id=d6d502c88f2cccf7 ip=127.0.0.1 udp=30399 tcp=30399
...
INFO [12-19|19:14:28.706] Starting Swarm HTTP proxy                port=8500
```

看到上面 INFO 的输出结果，说明 swarm 已经正常配置并运行，默认运行在 `8500` 端口上。

## 上传文件到 swarm

```go
package main

import (
    "fmt"
    "log"

    bzzclient "github.com/ethersphere/swarm/api/client"
)

func main() {
    client := bzzclient.NewClient("http://127.0.0.1:8500")

    file, err := bzzclient.Open("hello.txt")
    if err != nil {
        log.Fatal(err)
    }

    manifestHash, err := client.Upload(file, "", false, false, false)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(manifestHash) // 2e0849490b62e706a5f1cb8e7219db7b01677f2a859bac4b5f522afd2a5f02c0
}
```

然后就可以通过地址进行访问： `bzz://2e0849490b62e706a5f1cb8e7219db7b01677f2a859bac4b5f522afd2a5f02c0`

## 从swarm下载文件