# 智能合约

## 简介
为了可以和只能合约交互，必须首先生成合约的ABI(Application Binay interface)
再将ABI编译成可以导入到go中使用的格式

##  安装 slidity 编译器

```bash
	brew update
	brew tap ethereum/ethereum
	brew install solidity

	# 或者 

	npm install -g solc
```

## 安装 abigen 工具

该工具主要用于将 solidity 生成为 ABI

方式一:(推荐)

```bash
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
```

方式二:

```bash
go get -u github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make
make devtools
```

## 创建智能合约

为了便于测试，我们创建一个简单的合约

> 关于智能合约的编写推荐使用 truffle framework

```solidity
pragma solidity ^0.8.28;

contract Store {
  event ItemSet(bytes32 key, bytes32 value);

  string public version;
  mapping (bytes32 => bytes32) public items;

  constructor(string memory _version) {
    version = _version;
  }

  function setItem(bytes32 key, bytes32 value) external {
    items[key] = value;
    emit ItemSet(key, value);
  }
}
```

创建 go 合约文件

通过 abi 生成 

```bash
solc --abi Store.sol -o build
```

转换abi为go文件

```bash
abigen --abi=./build/Store.abi --pkg=store --out=store.go
```

> 生成的go文件里包含和智能合约交互的方法

为了部署智能合约，需要编译智能合约为 EVM 字节码， EVM 字节码是包含在交易的 `data` 字段里

```bash
solc --bin Store.sol -o build
```

以上操作也可以一次性生成

```bash
abigen --bin=./build/Store.bin --abi=./build/Store.abi --pkg=contract --out=store.go
```

## 部署智能合约

contract_deploy.go

## 加载智能合约

contract_load.go

## 查询智能合约

contract_read.go

## 写一个智能合约

contract_write.go

## 读取智能合约字节码

contract_bytecode.go

## 查询一个智能合约ERC20 token的智能合约

创建 ERC20 token

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

abstract contract ERC20 {
    string public constant name = "";
    string public constant symbol = "";
    uint8 public constant decimals = 0;

    function totalSupply() public view virtual returns (uint);
    function balanceOf(address tokenOwner) public view virtual returns (uint balance);
    function allowance(address tokenOwner, address spender) public view virtual returns (uint remaining);
    function transfer(address to, uint tokens) public virtual returns (bool success);
    function approve(address spender, uint tokens) public virtual returns (bool success);
    function transferFrom(address from, address to, uint tokens) public virtual returns (bool success);

    event Transfer(address indexed from, address indexed to, uint tokens);
    event Approval(address indexed tokenOwner, address indexed spender, uint tokens);
}
```

生成abi(json格式)

```bash
solc --abi erc20.sol -o build
```

生成go文件

```bash
abigen --abi=./build/ERC20.abi --pkg=token --out=token/erc20.go
```
