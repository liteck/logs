# logs
自用的简单日志记录,由于目前的代码都是通过 supervsion 部署的.自带了日志保存到文件和文件切割的功能,所以这里就不做处理了


### 使用方法
```go

logs.Debug("hello")
```

### 更新记录
```go

1. 独立日志包,使用系统log 库来做简单封装
2. 提供默认等级.并进行颜色配置
```