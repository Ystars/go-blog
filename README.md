## 技术选型
1. web:[gin](https://github.com/gin-gonic/gin)
2. orm:[gorm](https://github.com/jinzhu/gorm)
3. database:[mysql](https://github.com/go-gorm/mysql)
4. 权限配置:[casbin](https://github.com/casbin/casbin)
5. 跨域:[Gin-Cors](https://github.com/gin-contrib/cors): Gin框架提供的跨域中间件
6. 配置文件 [env](https://github.com/joho/godotenv)

## Godotenv

项目在启动的时候依赖以下环境变量，但是在也可以在项目根目录创建.env文件设置环境变量便于使用(建议开发环境使用)

```shell
MYSQL_HOST="127.0.0.1"      # MYSQL连接
MYSQL_PORT="3306"           # MYSQL端口
MYSQL_NAME="goblog"         # MYSQL数据库名
MYSQL_USER="root"           # MYSQL连接用户
MYSQL_PASSWORD="root"       # MYSQL连接密码

GIN_MODE="debug"
HTTP_PORT=":3030"

JWT_KEY="JweRaYTHIK"        # JWT密钥，必须设置而且不要泄露
```

## Go Mod

本项目使用[Go Mod](https://github.com/golang/go/wiki/Modules)管理依赖。

```shell
$ go mod tidy
$ go env -w GOPROXY=http://mirrors.aliyun.com/goproxy/
$ go run main.go // 自动安装
```

## 运行

```shell
$ go run main.go
```
