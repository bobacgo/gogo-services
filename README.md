# gogo-services

### 项目组件依赖
- polaris -- 注册中心、配置中心
- gorm -- orm 框架，操作关系型数据库
- 


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