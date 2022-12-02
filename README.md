# gogo-services

### 项目组件依赖

- polaris -- dns 注册中心、配置中心
- gin -- web framework
- gorm -- orm framework，操作关系型数据库
- go-redis -- redis

```text
gogo-services
    ├─admin-service -- 平台运营服务、用户中心
    ├─common-lib    -- 基础库
    ├─devops-conf   -- 部署相关
    ├─gateway       -- 后端网关
    └─main-service  -- 主服务
```

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

部署文档参考 devops-conf 模块

polaris、mysql、redis

#### hosts

```shell
127.0.0.1 polaris-headless
```