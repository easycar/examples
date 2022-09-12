# examples

虽然还未开发easycar的客户端，但是目前也是支持不需要客户端也可以使用的

#### withoutClient

有三个服务 order，account以及stock三个服务。

从协议上区分:account提供的是http服务。而order和stock提供的是grpc。

从事务模式上区分:account和stock提供的是TCC事务模式。 而order提供的是Saga事务模式。

由这三个服务组成一个分布式事务。
当用户下单时，需要经过这三个服务中内部一些接口。

如果只是同步执行第一阶段，那么第一阶段总执行时间 (account1+stock1+order1)。

很多场景下，分布式事务之间并不会存在直接的执行依赖先后关系。所以多个子事务一阶段可以同时并发执行。

流程就像这样：

![global](https://cdn.syst.top/servers.png)

假设当前我们需要保证创建订单前必须先执行成功account扣减余额和stock扣减库存数量，才能请求创建订单order的服务。同时account和stock服务并不需要保证他们的执行顺序。

那么我们一阶段总执行耗时可以粗略=max(account1,stock1)+order1。

##### start

```shell
go run withoutclient/main.go base ## 目前demo只提供上述组合一种场景，后续继续细化更多demo
```


#### client

To be developed