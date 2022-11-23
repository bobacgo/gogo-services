# gogo-services

### create project
> 守护进程的项目以 -service 结尾
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