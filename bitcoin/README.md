# Bitcoin Core

## 准备工作

```bash
# 构建要求
$ sudo apt-get install build-essential libtool autotools-dev automake pkg-config bsdmainutils python3
# 安装所需的依赖项
$ sudo apt-get install libevent-dev libboost-dev
# 钱包需要 SQLite
$ sudo apt install libsqlite3-dev
# 可选的端口映射库
$ sudo apt install libminiupnpc-dev libnatpmp-dev
# ZMQ依赖(提供ZMQ API)
$ sudo apt-get install libzmq3-dev
# 用户空间、静态定义的跟踪 (USDT) 依赖项
$ sudo apt install systemtap-sdt-dev
# 要使用 Qt 5 进行构建, 您需要以下内容
$ sudo apt-get install libqt5gui5 libqt5core5a libqt5dbus5 qttools5-dev qttools5-dev-tools
# 此外, 为现代桌面环境支持 Wayland 协议
$ sudo apt install qtwayland5
# libqrencode (可选) 可以安装
$ sudo apt-get install libqrencode-dev
```

## 从源码编译 Bitcoin Core

```bash
# 下载源码
$ git clone https://github.com/bitcoin/bitcoin.git
# 选择 Bitcoin Core 版本
$ git tag
$ git checkout v25.0
# 配置构建 Bitcoin Core
$ ./autogen.sh
$ ./configure
# 构建 Bitcoin Core 可执行文件
$ make
$ make check && sudo make install
```

## 运行 Bitcoin Core 节点

- 配置 Bitcoin Core 全节点

[import, title:"~/.bitcoin/bitcoin.conf", lang:"text"](bitcoin.conf)

```bash
# 运行 主网 节点
$ bitcoind
# 运行 Testnet 节点
$ bitcoind -testnet
```

bitcoin.conf 配置文件
除了 -conf 以外的所有命令行参数都可以通过一个配置文件来设置.
请注意，如果在命令行和配置文件中同时设置某个参数，命令行中的设置将覆盖配置文件中的设置。

# 网络相关的设置：

# 通过一个 Socks4 代理服务器连接
#proxy=127.0.0.1:1080

##############################################################
##            addnode 与 connect 的区别                     ##
##                                                          ##
##  假设您使用了 addnode=4.2.2.4 参数，那么 addnode 便会与   ##
##  您的节点连接，并且告知您的节点所有与它相连接的其它节点。   ##
##  另外它还会将您的节点信息告知与其相连接的其它节点，这样它   ##
##  们也可以连接到您的节点。                                 ##
##                                                          ##
##  connect 在您的节点“连接”到它的时候并不会做上述工作。仅  ##
##  它会与您连接，而其它节点不会。                           ##
##                                                          ##
##  因此如果您位于防火墙后，或者因为其它原因无法找到节点，则   ##
##  使用“addnode”添加一些节点。                            ##
##                                                          ##
##  如果您想保证隐私，使用“connect”连接到那些您可以“信任” ##
##  的节点。                                                ##
##                                                          ##
##  如果您在一个局域网内运行了多个节点，您不需要让它们建立许多 ##
##  连接。您只需要使用“connect”让它们统一连接到一个已端口转  ##
##  发并拥有多个连接的节点。                                 ##
##############################################################
# 您可以使用多个 addnode= 设置来连接到指定的节点
# 也可以使用多个 connect= 设置来仅连接到指定的节点

# 不使用因特网中继聊天(IRC)频道来查找其它节点
#noirc=0

# 入站+出站的最大连接数
#maxconnections=

# 客户端在 HTTP 连接建立后，等待多少秒以完成一个 RPC HTTP 请求
#rpctimeout=30

# 可以使用 * 作为通配符
#rpcallowip=192.168.1.*

# 使用安全套接层(也称为 TLS 或 HTTPS)来连接到 Bitcoin -server 或 bitcoind
#rpcssl=1

# 当 rpcssl=1 时使用的 OpenSSL 设置
#rpcsslciphers=TLSv1+HIGH:!SSLv2:!aNULL:!eNULL:!AH:!3DES:@STRENGTH
#rpcsslcertificatechainfile=server.cert
#rpcsslprivatekeyfile=server.pem 

# 设置 gen=1 以尝试生成比特币(采矿)
#gen=0

# 预生成如下数目的公钥和私钥, 这样钱包备份便可以对已有的交易以及未来多笔交易有效
#keypool=100

# 每次发送比特币的时候支付一个可选的额外的交易手续费. 
# 包含手续费的交易会更快的被包含在新生成的货币块中, 因此会更快生效
#paytxfee=0.00

# 允许直接连接，实现“通过 IP 地址支付”功能
#allowreceivebyip=1

[获取测试币](https://coinfaucet.eu/en/btc-testnet/)

执行一次挖矿操作
bitcoin-cli -testnet -generate 1

[创建本地比特币网络regtest，手工创建交易，极速理解比特币原理](https://zhuanlan.zhihu.com/p/592586339)

[精通比特币](https://github.com/inoutcode/bitcoin_book_2nd/tree/master)