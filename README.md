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