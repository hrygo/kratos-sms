# Kratos-sms

## Quick Start

1、安装kratos

```bash
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
kratos upgrade
```

2、更新下载依赖

```bash
cd kratos-sms
go mod download
go mod tidy
```

3、安装环境依赖
安装 docker、docker-compose、consul

4、重新生成代码保证代码最新

```bash
make api
make config
make generate
```

5、运行依赖环境

```
make set-env
```

6、调试执行应用

```bash
# 直接执行
go run cmd/kratos-sms/main.go
# 或者执行
kratos run
```

7、构建可执行程序并运行

```bash
make build

cd bin
./kratos-sms --conf ./configs
```

参考 [Kratos Layout](./Kratos_layout.md) 官方文档


