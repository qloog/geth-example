# 客户端

本地启用测试网络，方便进场本地测试

```bash
# 安装
npm install -g ganache-cli

# 启动测试网服务
ganache-cli

# 测试
cd client
go run client.go
# Output: 
#   we have a connection
```

成功启动后

```bash
➜  go-eth git:(main) ✗ ganache-cli
Ganache CLI v6.12.2 (ganache-core: 2.13.2)
(node:76551) [DEP0040] DeprecationWarning: The `punycode` module is deprecated. Please use a userland alternative instead.
(Use `node --trace-deprecation ...` to show where the warning was created)

Available Accounts
==================
(0) 0xEF08f8703820D85BaAab554222576C4C369732f1 (100 ETH)
(1) 0x261Fe4CD41d6ee229581258046bE3CB5Dac7EA7e (100 ETH)
(2) 0x9dC4Cd46e71C7f43d0656B70134b868979d48497 (100 ETH)
(3) 0x780bE64Ab243E90E70D1c17A86f56F3aA63BF891 (100 ETH)
(4) 0x5e429A6d9ceB389a90538A825b8b17De853dE82d (100 ETH)
(5) 0x468B1C7C15770210fc13D0Ef85522DC3CeDac7dA (100 ETH)
(6) 0xB285f66e53B6cD7842D572C8dc91B87B3f0A5f1d (100 ETH)
(7) 0xf90B26F90C8C2c5eE5ED1EE25a860F1686493Db6 (100 ETH)
(8) 0x1069CA0BE8882493baab1DC072f0606D272A9043 (100 ETH)
(9) 0xafBAb784859bC501F23DB227b0431f3DAaDB1000 (100 ETH)

Private Keys
==================
(0) 0xb71f694a79a7ed74ce604467b0183019adc09f84ebd1ff9f79c1ed0542451321
(1) 0x7586fa7b35cd708262fb1c02662c8c97540437c6705aa60b239c66e422dbfbd8
(2) 0xf366a7a46ac34f1746e3e8ec9446900f4cd3b4f6f78c99d4f74924aae1d86280
(3) 0xc140c7f713adef2845cdc57b25ec582aeb80721c37e2b8087c81bc1e37df8a88
(4) 0xcdaa0573b5566fb3dd02f94db5ee169dd625a707b2817d792bf22f9e2a756314
(5) 0x39647bafefa53ddd360280d0f05e50f9562117e4e3a3e33382b70272c6e273fc
(6) 0x17a14928ef15afb12825217e3c2ecf738c85981810b3c5fdefa66afe91baa7df
(7) 0xae1f571b9934a5e39442862a05b7eecbab30ad76dd9b1865870ba9fe12cc1e0e
(8) 0x1afe6881a7375a488f3e7d8a416339e9546c1b2fbd45eda282995f10ce311853
(9) 0x7b7f0260fad88eaf66d1e9ac74ced876b33b520351a697d902a007b2105470ee

HD Wallet
==================
Mnemonic:      inch elite client hospital candy arm suspect fiber party visual glare spot
Base HD Path:  m/44'/60'/0'/0/{account_index}

Gas Price
==================
20000000000

Gas Limit
==================
6721975

Call Gas Limit
==================
9007199254740991

Listening on 127.0.0.1:8545
```

