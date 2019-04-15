# DolphinChain

![dolphinchain.logo](./doc/images/dolphinchain.logo.png)

`DolphinChain` 是世界上第一个区块链应用靶机！[dolphinchain.org](http://dolphinchain.org)

Version : 1.0.0

## Table of Contents

<!-- TOC -->

- [DolphinChain](#dolphinchain)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Tendermint Bugs History](#tendermint-bugs-history)
  - [Contribution](#contribution)
  - [Backer](#backer)
  - [Connection](#connection)
  - [TODO](#todo)
  - [License](#license)

<!-- /TOC -->

## Overview

`DolphinChain` 是由[玄猫安全实验室](https://github.com/XuanMaoSecLab)维护的区块链应用靶机，旨在教授区块链应用程序安全课程。您可以使用 `DolphinChain` 进行安装和练习。

`DolphinChain` 基于 `tendermint v0.31.2 (WARNING: ALPHA SOFTWARE)` 开发，是当时的 `tendermint` 最新版本。

在这个版本里(v1.0.0)，我们在`DolphinChain`设置了10多个缺陷。任何白帽子与区块链开发者都可以尝试挖掘漏洞。`DolphinChain`目的在于帮助安全人员提高技能，同时帮助区块链开发者更好地了解保护区块链应用程序的过程。

## Installation

1. 下载并安装 [golang](https://golang.org/doc/install)

2. 获取 `DolphinChain`.

3. 获取 `DolphinChain` 的所有依赖.

只要以上步骤，你就已经准备好。

你可以在 [部署文档](./doc/Installation.md) 找到更详细的安装步骤。

## Usage

部署 `DolphinChain` -> 寻找缺陷代码 -> 编写验证脚本 -> 验证漏洞存在

1. 寻找漏洞 ：区块链漏洞主要由代码问题及业务逻辑问题引起。
2. 编写验证脚本 ：有两种方式，PoC 或  Go test 测试脚本。

当然我们会公开所有漏洞 `Writeup` 。您可以通过我们另一个仓库查看。

同时，我们也归纳了 `tendermint` 的历史漏洞，见 [Tendermint Bugs History](#Tendermint-Bugs-History)

## Tendermint Bugs History

tendermint 是 Cosmos 网络生态体系中的核心组件，主要负责共识与P2P。从2014年开发至今，社区活跃高、代码迭代快，最重要的是，非常重视安全性。因此我们通过学习此链的安全漏洞及修复方案，可以让其它开发者学习思路及避免踩到前人已经踩过的坑。

这里归纳了 tendermint 的历史漏洞，这我们花费了近一个月时间整理。

|   |  P2P  |  共识  |  节点  |  接口  |  序列化反序列化  |  消息队列  |  数据库  |  消息  |  链逻辑  |  种子列表  |  内存池  |
|----|----|----|----|----|----|----|----|----|----|----|----|
空指针  |  X  |  X  |  X  |    |  X  |    |  X  |  X  |  X  |  X  |  X  |
配置为空  |    |  X  |  X  |    |  X  |    |  X  |  X  |  X  |  X  |  X  |
缺乏异常处理  |  X  |  X  |  X  |  X  |  X  |  X  |  X  |  X  |  X  |    |  X  |
服务挂起  |    |    |    |    |    |    |    |    |    |    |  X  |
并发数量限制  |  X  |    |  X  |  X  |  X  |    |    |  X  |    |    |  X  |
畸形数值  |  X  |  X  |    |  X  |  X  |    |  X  |  X  |    |  X  |    |
组件逻辑  |  X  |  X  |    |    |    |  X  |    |    |  X  |    |    |
溢出  |    |    |    |    |  X  |  X  |    |    |    |    |    |
死锁(无随机性)  |  X  |    |  X  |    |    |    |    |    |    |    |    |
DOS  |  X  |    |    |    |    |    |  X  |    |    |  X  |  X  |
内存泄露导致OOM  |  X  |    |    |  X  |    |    |    |  X  |    |    |    |
初始化  |  X  |  X  |  X  |    |    |  X  |  X  |    |  X  |    |  X  |
程序依赖  |    |    |  X  |    |    |    |    |  X  |    |    |    |
资源控制  |  X  |    |  X  |    |    |  X  |    |  X  |    |    |  X  |

## Contribution

欢迎通过 issue 提交问题。同时，您也可以跟我们一起开发更多的漏洞。

贡献者：

Tri0nes、Javierlev

## Backer

<p>
  <a href="https://www.bugx.io" target="_blank"><img src="./doc/images/bugx.png" width="150"></a>
  <br>
  <br>
  <a href="https://github.com/XuanMaoSecLab" target="_blank"><img src="./doc/images/xuanmao_logo.jpg" width="150"></a>
  <br><br>
  <img src="./doc/images/blockseccode.jpg" width="150">
</p>

## Connection

<p>
<img src="./doc/images/Tri0nes.jpg" width="200">
</p>

## TODO

- [ ] 链上还有些bugs我们正在修补。
- [ ] 针对已有的约10个漏洞编写`Writup`
- [ ] 梳理新的漏洞作为后续漏洞开发
- [ ] 一些特别有趣的想法

## License

DolphinChain is licensed under the MIT License. See [LICENSE](./LICENSE) for the full license text.
