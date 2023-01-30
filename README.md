## go-zero-lol

`go-zero-lol` 是基于golang微服务框架 [go-zero](https://github.com/zeromicro/go-zero) 开发的单体应用项目，api直连mysql数据库，没有使用rpc。

集成组件：

1. 支持 [tokenlimit](https://go-zero.dev/cn/docs/blog/governance/tokenlimit) 令牌桶限流 
1. 支持 [middleware](https://go-zero.dev/cn/docs/advance/middleware) 中间件使用 
1. 支持 [jwt](https://go-zero.dev/cn/docs/advance/jwt) 鉴权 
1. 支持 [Prometheus](https://github.com/prometheus/client_golang) 指标记录 
1. 支持 [Swagger](https://github.com/swaggo/gin-swagger) 接口文档生成 
1. 支持 trace 项目内部链路追踪 
1. 支持 [pprof](https://github.com/gin-contrib/pprof) 性能剖析
1. 支持 [errorx](https://go-zero.dev/cn/docs/advance/error-handle) 统一定义错误码 
1. 支持 [zap](https://go.uber.org/zap) 日志收集 
1. 支持 [go-redis](https://github.com/go-redis/redis/v7) 组件
1. 支持 RESTful API 返回值规范
1. 支持 [rpc编写与调用](https://go-zero.dev/cn/docs/advance/rpc-call)
1. 支持 生成数据表 [CURD](https://go-zero.dev/cn/docs/advance/model-gen)、[控制器方法](https://go-zero.dev/cn/docs/goctl/api) 等代码生成器

## 目录说明

```text
.
├── Dockerfile
├── LICENSE
├── README.md
├── api ---------------------------------------- api定义文件目录
│   ├── download.api
│   ├── login.api
│   ├── lol.api -------------------------------- lol.api为整个项目api的定义文件，import其他模块api文件
│   └── upload.api
├── etc
│   └── lol-api.yaml --------------------------- 项目配置文件
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   ├── handler
│   │   ├── downloadhandler.go
│   │   ├── loginhandler.go
│   │   ├── routes.go
│   │   └── uploadhandler.go
│   ├── logic
│   │   ├── downloadlogic.go
│   │   ├── loginlogic.go
│   │   └── uploadlogic.go
│   ├── middleware ----------------------------- 中间件目录
│   │   ├── checkpathmiddleware.go
│   │   └── tokenlimitermiddleware.go
│   ├── svc
│   │   └── servicecontext.go
│   └── types
│       └── types.go
├── lol.go 
└── static ------------------------------------- 静态资源文件目录
```

## 常用命令

- goctl一键安装protoc & protoc-gen-go
```
goctl env check -i -f --verbose
```

- goctl创建api模板
```
cd go-zero-lol/
goctl api -o lol.api
```

- goctl生成api服务
```
cd go-zero-lol/
goctl api go -api lol.api -dir .
```

- goctl创建rpc模板
```
cd rpc/
goctl rpc -o=user.proto
```

- gotcl生成rpc服务
```
cd rpc/
goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
```

- gotcl生成model
```
cd model/
goctl model mysql ddl -src user.sql -dir . -c
```

- gotcl生成Dockerfile
```
cd go-zero-lol/
goctl docker -go lol.go
```

## 其他

- docker单机部署etcd
```
docker run -d --name etcd-server \
    --publish 2379:2379 \
    --publish 2380:2380 \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    --env LISTEN-CLIENT-URLS=http://0.0.0.0:2379 \
    bitnami/etcd:latest
```

