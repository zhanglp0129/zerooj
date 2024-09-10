# judge
测评相关文档，主要包含测评的流程

## 流程
1. 用户点击提交按钮，前端将用户id、题目id、编程语言、代码等信息发送到网关服务，并建立长连接
2. 网关服务调用题库服务，将提交记录写入数据库，测评结果为等待中，并发送到前端
3. 网关服务将题目id、代码等评测需要的数据发送到测评服务，将测评结果改为测评中，并发送到前端
4. 测评服务获取测评数据后，开始测评
5. 测评结束后，测评服务返回测评结果；网关服务将结果写入数据库，并发送到前端
6. 结束测评，关闭长连接

## 服务注册etcd的key命名
- 一个测评服务进程只能处理一个语言。通过etcd的key来指定处理的语言，并实现负载均衡
- 一台服务器或docker容器内可以启动多个测评服务进程。语言的版本包含在语言名称中
- key命名格式：judge.rpc.语言名称，不忽略大小写

|    语言名称    |         key          |
|:----------:|:--------------------:|
|    C++     |    judge.rpc.C++     |
|     C      |     judge.rpc.C      |
|     C#     |     judge.rpc.C#     |
|    Java    |    judge.rpc.Java    |
|  Python2   |  judge.rpc.Python2   |
|  Python3   |  judge.rpc.Python3   |
| JavaScript | judge.rpc.JavaScript |
| TypeScript | judge.rpc.TypeScript |
|     Go     |     judge.rpc.Go     |
|    Rust    |    judge.rpc.Rust    |
|    Ruby    |    judge.rpc.Ruby    |
|    PHP     |    judge.rpc.PHP     |
|   Swift    |   judge.rpc.Swift    |


## 注意
- 当调用测评服务返回error时，测评结果改为系统错误