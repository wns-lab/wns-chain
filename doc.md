# WNS protocol
wns protocol 是一个web3原生的跨链DID协议，它旨在通过跨链架构为web3原生dapp以及各种web2应用提供统一便捷的域名服务。

# wns
web3域名是wns的具体表现形式，将用户在多条链上的账户地址和名称映射到以`.web3`为后缀的域名，实现多链账户地址和名称的统一管理和便捷操作

wns架构分为两部分：
- wns chain
    - wns本身是一个分布在多链上的应用，为了整合在多链上的信息，并且保证多链一致的全局状态，我们开发一条cosmos application chain作为全局状态的维护者，保证各条链的状态不冲突并负责不同链之间的信息传递的relay工作。
    - 概念：
        - cosmos application chain
            - 与常见的有vm的链不同的是 cosmos application chain鼓励dapp搭建有自我主权的应用链，其中tendermint负责共识层和网络层，项目方使用cosmos-sdk进行应用层的开发。应用层的逻辑以module的形式被组合在一起，每个module拥有独立的store，以及该module特定的message type和message server，以及相应的message的处理逻辑。当一笔交易被应用处理时，交易内的message根据不同的type被路由到不同的message server被处理。module之间可以通过grpc通信。同时cosmos应用链还可以在module的基础上支持evm执行合约代码。
            - 已有的module：常规的auth, bank, staking等，group module可以用来实现DAO治理等功能
            - 我们需要实现的module：将ENS的功能移植成为wns module，完成web3域名的注册管理等功能。另外需要探索实现ERC4377账户抽象
            - 第一阶段将用户注册web3域名的入口放在wns应用链上，由wns应用链负责将注册信息同步到对应的原生公链。
            - 在wns应用链上完成DAO的治理，share分红等逻辑
            - 参考文档和代码仓库：
                - https://docs.cosmos.network/main
                - https://github.com/cosmos/cosmos-sdk
- wns-contracts
    - 在除了wns应用链之外的公链上，wns将部署智能合约管理该链上对应的web3域名的部分状态
    - wns合约通过由wns应用链跨链维护相应的状态以支持合约所在公链其他合约调用域名解析功能，核心是`gravity.sol`合约,`Registry.sol`合约和`PublicResolver.sol`合约
- gravity bridge:
    - 为了支持wns应用链跨链同步注册信息，我们需要集成gravity bridge
    - gravity bridge 原理: wns应用链的验证节点构成wns应用链验证人集合，以及对应的evm上gravity.sol的多签钱包集合，通过收集超过2/3验证节点签名的方式保证跨链消息的准确性
        - gravity.sol => 每个wns节点需运行一个eth full node，用于捕获gravity.sol产生的event并出发对应的wns应用链的功能
        - wns 应用链 => gravity.sol: 消息被wns验证节点签名后relay到gravity.sol，gravity.sol需验证2/3多签。
    - 参考资料以及代码仓库：
        - https://blog.althea.net/how-gravity-works/
        - https://github.com/Gravity-Bridge/Gravity-Bridge
        - https://github.com/Gravity-Bridge/Gravity-Bridge