# Go Shopping
This example contains a suite of microservices all built on the `go micro` framework. The goal of this example is to provide you
with a practical example of multiple services operating within a lightweight ecosystem that gives you more than just
the simple "hello world" functionality. In addition to 3 backing services, there is a single aggregating API service that mimicks 
and extremely common best practice in cloud native/microservice development.

### 编译
```
make all 
```

### 运行

#### Redis

```shell
# 安装和启动Redis
sudo apt-get install redis-server
sudo /etc/init.d/redis-server  start
```

#### consul 

```shell
# 注册中心
./consul agent -dev
```

#### catalogd

```sh
go run catalog/cmd/catalogd/main.go
```

#### shippingd

```shell
go run shipping/cmd/shippingd/main.go
```

#### warehoused

```shell
go run warehouse/cmd/warehoused/main.go
```

#### api

```shell
go run api/cmd/apid/main.go
```



