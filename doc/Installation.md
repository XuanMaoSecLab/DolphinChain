# DolphinChain安装部署

## 部署

`DolphinChain`主要基于`tendermint`进行开发。安装步骤与`tendermint`类似。由于现阶段主要目的并不是运行起一条主链，而是通过测试、代码审计来发现潜在的漏洞，所以主要介绍必要的依赖安装。

## 依赖安装

### 安装dep工具

```bash
go get github.com/golang/dep
cd $GOPATH/src/github.com/golang/dep
go install ./...
vim /etc/profile
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

### 手动安装基本依赖

(以下依赖不一定完全，根据实际运行测试运行时需要再装相关依赖也是可以的)：

创建一个脚本文件`dep.sh`

```bash
go get github.com/btcsuite/btcd/btcec
go get github.com/davecgh/go-spew/spew
go get github.com/go-kit/kit/log
go get github.com/go-kit/kit/log/level
go get github.com/go-kit/kit/log/term
go get github.com/go-kit/kit/metrics
go get github.com/go-kit/kit/metrics/discard
go get github.com/go-kit/kit/metrics/prometheus
go get github.com/go-logfmt/logfmt
go get github.com/golang/protobuf/proto
go get github.com/golang/protobuf/ptypes
go get github.com/golang/protobuf/ptypes/any
go get github.com/golang/protobuf/ptypes/duration
go get github.com/golang/protobuf/ptypes/timestamp
go get github.com/gorilla/websocket
go get github.com/pkg/errors
go get github.com/prometheus/client_golang/prometheus
go get github.com/prometheus/client_golang/prometheus/promhttp
go get github.com/rcrowley/go-metrics
go get github.com/rs/cors
go get github.com/spf13/cobra
go get github.com/spf13/viper
go get github.com/syndtr/goleveldb/leveldb
go get github.com/syndtr/goleveldb/leveldb/errors
go get github.com/syndtr/goleveldb/leveldb/iterator
go get github.com/syndtr/goleveldb/leveldb/opt
go get github.com/gogo/protobuf/gogoproto
go get github.com/gogo/protobuf/jsonpb
go get github.com/gogo/protobuf/proto
go get github.com/gogo/protobuf/types
go get github.com/tendermint/go-amino
```

### 安装特殊依赖

对于一些不能直接获取的依赖，采用间接获取的方法

#### Linux下

Linux下创建脚本`dep2.sh`

```sh
mkdir $GOPATH/src/golang.org/x/
git clone https://github.com/golang/crypto.git $GOPATH/src/golang.org/x/crypto
git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net
git clone https://github.com/golang/sys.git $GOPATH/src/golang.org/x/sys
git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text
git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc
git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto
git clone https://github.com/go-yaml/yaml.git $GOPATH/src/gopkg.in/yaml.v2
```

### windows下

windows下`GOPATH`环境变量应为`%GOPATH%`，所以把以上`$GOPATH`改成`%GOPATH%`即可

windows下创建批处理文件`dep2.bat`，然后再终端运行

```sh
git clone https://github.com/golang/crypto.git %GOPATH%/src/golang.org/x/crypto
git clone https://github.com/golang/net.git %GOPATH%/src/golang.org/x/net
git clone https://github.com/golang/sys.git %GOPATH%/src/golang.org/x/sys
git clone https://github.com/golang/text.git %GOPATH%/src/golang.org/x/text
git clone https://github.com/grpc/grpc-go.git %GOPATH%/src/google.golang.org/grpc
git clone https://github.com/google/go-genproto.git %GOPATH%/src/google.golang.org/genproto
git clone https://github.com/go-yaml/yaml.git %GOPATH%/src/gopkg.in/yaml.v2
```

### 测试工具相关依赖

在测试文件时会用到

```sh
go get github.com/stretchr/testify/require
```

## 编译

编译：

```bash
make install
make build
```

build后的文件位于目录/build下
能够执行 `build/tendermint`，说明安装成功：

```bash
root@ubuntu:~/build# tendermint
Tendermint Core (BFT Consensus) in Go

Usage:
  tendermint [command]

Available Commands:
  gen_node_key                Generate a node key for this node and print its ID
  gen_validator               Generate new validator keypair
  help                        Help about any command
  init                        Initialize Tendermint
  lite                        Run lite-client proxy server, verifying tendermint rpc
  node                        Run the tendermint node
  probe_upnp                  Test UPnP functionality
  replay                      Replay messages from WAL
  replay_console              Replay messages from WAL in a console
  show_node_id                Show this node's ID
  show_validator              Show this node's validator info
  testnet                     Initialize files for a Tendermint testnet
  unsafe_reset_all            (unsafe) Remove all the data and WAL, reset this node's validator to genesis state
  unsafe_reset_priv_validator (unsafe) Reset this node's validator to genesis state
  version                     Show version info

Flags:
  -h, --help               help for tendermint
      --home string        directory for config and data (default "/root/.tendermint")
      --log_level string   Log level (default "main:info,state:info,*:error")
      --trace              print out full stack trace on errors

Use "tendermint [command] --help" for more information about a command.
```

## 运行

```bash
tendermint init
tendermint node --proxy_app=kvstore
# export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

```bash
root@ubuntu:~/# netstat -tnlp
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name
tcp6       0      0 :::26656                :::*                    LISTEN      9229/tendermint 
tcp6       0      0 :::26657                :::*                    LISTEN      9229/tendermint 
```

## RPC请求

```bash
curl -s localhost:26657/status
```

回显如下:

```bash
Available endpoints:
//localhost:26657/abci_info
//localhost:26657/consensus_state
//localhost:26657/dump_consensus_state
//localhost:26657/genesis
//localhost:26657/health
//localhost:26657/net_info
//localhost:26657/num_unconfirmed_txs
//localhost:26657/status

Endpoints that require arguments:
//localhost:26657/abci_query?path=_&data=_&height=_&prove=_
//localhost:26657/block?height=_
//localhost:26657/block_results?height=_
//localhost:26657/blockchain?minHeight=_&maxHeight=_
//localhost:26657/broadcast_tx_async?tx=_
//localhost:26657/broadcast_tx_commit?tx=_
//localhost:26657/broadcast_tx_sync?tx=_
//localhost:26657/commit?height=_
//localhost:26657/consensus_params?height=_
//localhost:26657/subscribe?query=_
//localhost:26657/tx?hash=_&prove=_
//localhost:26657/tx_search?query=_&prove=_&page=_&per_page=_
//localhost:26657/unconfirmed_txs?limit=_
//localhost:26657/unsubscribe?query=_
//localhost:26657/unsubscribe_all?
//localhost:26657/validators?height=_
```