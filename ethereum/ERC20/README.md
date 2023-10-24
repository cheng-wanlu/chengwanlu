## ERC20 智能合约信息

流程如下：

Solidity(ERC20.sol) –(solc)–> ABI(ERC20.abi) –(abigen)–> Go Package(erc20.go)

```sh
# 下载 solc
wget https://github.com/ethereum/solidity/releases/download/v0.8.21/solc-static-linux
chmod +x solc-static-linux
mv solc-static-linux /usr/bin/solc

# 下载 ERC20.sol
git clone https://github.com/OpenZeppelin/openzeppelin-contracts.git
# 构建 abi
solc --abi ~/src/OpenZeppelin/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol --base-path ~/src/OpenZeppelin/openzeppelin-contracts --output-dir ./

# 安装 abigen
git clone https://github.com/ethereum/go-ethereum.git 
cd go-ethereum
make && make all
ll build/bin/

~/src/ethereum/go-ethereum/build/bin/abigen --abi=ERC20.abi --pkg=token --out=erc20.go
```

[import, lang:"golang"](erc20.go)