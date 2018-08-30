# gogs webhook 自动更新代码

## 安装 

使用 github.com/constabulary/gb  编译

或者 进入 src/cmd/main 编译

## 使用

./path/to/bin/main -c ./path/to/src/com/main/app.toml


访问 webhook 地址
http://localhost:port/hooks?k=<名称>

## 配置
使用 github.com/BurntSushi/toml

```
listen = ":8081"
<!--  这是仓库的名字-->
[Reps.test]
<!--  这是自己随便定义的名字-->
name = "test"
ref="refs/heads/master"
secret = "Secret"
srcPath = "/data/test"
<!--  可以多个用户-->
allowUser = ["weisd","fuck"]

[Reps.a]
name = "a"
ref="refs/heads/master"
secret = "Secret"
srcPath = "SrcPath"
allowUser = ["weisd"]

[Reps.b]
name = "b"
ref="refs/heads/master"
secret = "Secret"
srcPath = "/data/b"
allowUser = ["weisd"]
```


## 打包代码
```
➜  main git:(master)     GOPATH=/Users/dadadadada/Documents/GitHub/gogs-webhooks/vendor

➜  main git:(master)     CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gogs_web_hook main.go struct.go config.go

```

## linux 挂起命令，并且写入日志

```

nohup ./gogs_web_hook -c app.toml > hook.log 2>&1 &

```