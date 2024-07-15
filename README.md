# 网站主页

# 手动下载安装运行

## 1、clone 项目

## 2、复制 `config.example.yml` 为 `config.yml`，修改成自己的配置文件

## 3、下载依赖

```go
go mod tidy
```

## 4、启动项目

```go
go run main.go
```

# 使用 `Docker`

```docker
docker run --name=homepage -v config.yml:/src/config.yml -p 8080:8080 -d ghcr.io/xiaoxuan6/homepage:latest
```

