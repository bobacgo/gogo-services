# gogo-services

### 项目组件依赖

- polaris -- dns 注册中心、配置中心
- gin -- web framework
- gorm -- orm framework，操作关系型数据库
- go-redis -- redis

### create project

> 1. 守护进程的项目以 -service 结尾
> 2. 每个子项目都要有 README.md 文件

```shell
# first create
go work init
```

```shell
# add project
cd ./xxx-service
go mod init github.com/gogoclouds/gogo-services/xxx-service
cd ../
go work use ./xxx-service
```

### Develop Env

### 注册中心和配置中心

```shell
# 北极星 polaris
docker run --name polaris -p 45010:15010 -p 48080:8080 -p 48090:8090 -p 48091:8091 -p 48093:8093 -p 48761:8761 -p 49000:9000 -p 49090:9090 -d polarismesh/polaris-server-standalone:latest
```

### hosts

```shell
127.0.0.1 polaris-headless
```
