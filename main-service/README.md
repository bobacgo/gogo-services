# main-service

主服务、提供http接口、主要对接前端。

### Run

```shell
cd ./cmd
go run main.go -config ./config/polaris.yaml

# test 
# 注意：启动成功后服务会被注入提供健康监测的API
curl http://127.0.0.1:8000/health
```
