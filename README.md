# examples

## srv目录

有三个服务 order，account以及stock三个服务。

从协议上区分:account提供的是http服务。而order和stock提供的是grpc。

从事务模式上区分:account和stock提供的是TCC事务模式。 而order提供的是Saga事务模式。

当用户下单时，需要经过这三个服务中内部一些接口。

如果只是同步执行第一阶段，那么第一阶段总执行时间 (account1+stock1+order1)。

很多场景下，分布式事务之间并不会存在直接的执行依赖先后关系。所以多个子事务一阶段可以同时并发执行。

流程就像这样：

![global](https://cdn.syst.top/servers.png)

假设当前我们需要保证创建订单前必须先执行成功account扣减余额和stock扣减库存数量，才能请求创建订单order的服务。同时account和stock服务并不需要保证他们的执行顺序。

那么我们一阶段总执行耗时可以粗略=max(account1,stock1)+order1。

## start

确保已经启动了easycar服务端!!!

启动srv里的服务,

```shell
go run srv/main.go
```

client demo包括直连、服务发现、以及tls三种方式。

```shell
NAME:
   examples - examples for easycar

USAGE:
   examples [global options] command [command options] [arguments...]

COMMANDS:
   direct, direct        connection easrcar direct
   tls                   connection easrcar with tls
   discovery, discovery  connection easycar by discovery
   http, http            just request easycar service by http
   help, h               Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --easycar value  set easycar server url (default: "127.0.0.1:8089")
   --help, -h       show help (default: false)

```

**直连**

```shell
go run main.go --easycar="127.0.0.1:8089"  direct
```

**服务发现**

```shell
go run main.go --easycar="easycar" discovery
```

