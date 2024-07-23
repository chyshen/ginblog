# ginblog

[![MIT 许可证](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GitHub stars](https://img.shields.io/github/stars/chyshen/repo.svg?style=social&label=Star)](https://github.com/chyshen/ginblog/stargazers)
[![Gitee stars](https://gitee.com/chyshen/ginblog/badge/star.svg?theme=white)](https://gitee.com/chyshen/ginblog/stargazers)
[![Go ](https://img.shields.io/badge/Go-Package-blue.svg)]()

#### 介绍

个人博客系统

#### 目录结构


#### 软件架构

1.  前端（frontend）

    - Web前端框架：[Vue3](https://cn.vuejs.org/)
    - UI框架：[Element Plus](https://element-plus.org/zh-CN/)
    - Vue组件框架：[Vuetify](https://vuetifyjs.com/zh-Hans/)
    - 富文本编辑器：[TinyMCE](https://www.tiny.cloud/)，[TinyMCE中文文档](http://tinymce.ax-z.cn/)，[wangEditor](https://www.wangeditor.com/)（国产，中文文档齐全，移动端只能查看不能编辑）
    - 日期和时间格式化：[Day.js](https://day.js.org/docs/zh-CN/)（Element Plus使用）或 [Moment.js](https://momentjs.cn/)
    - 网络请求库：[Axios](https://www.axios-http.cn/)

2.  后端（backend）

    - Web框架：[Gin](https://gin-gonic.com/zh-cn/)
    - ORM框架：[GORM](https://gorm.io/zh_CN/)
    - Token：[jwt-go](https://golang-jwt.github.io/jwt/usage/create/)
    - 加密：[scrypt](https://pkg.go.dev/golang.org/x/crypto/scrypt)
    - 日志：[Logrus](https://pkg.go.dev/github.com/sirupsen/logrus)
    - 跨域：[gin-contrib/cors](https://github.com/gin-contrib/cors)
    - 配置文件：[Viper](https://github.com/spf13/viper) 或 [ini.v1](https://ini.unknwon.io/)
    - 接口文档：[Swagger](https://github.com/swaggo/swag)

3.  数据库

    - MySQL：[V8.0.21](https://downloads.mysql.com/archives/community/)


#### 安装教程

##### 克隆仓库

```shell
git clone https://gitee.com/chyshen/ginblog.git
# or
git clone https://github.com/chyshen/ginblog.git
```

##### 前端（frontend）

1.  安装

```shell
cd /ginblog/web
pnpm init
```

2.  运行

```shell
pnpm dev
```

##### 后端（backend）

1.  安装

```shell
cd /ginblog/api
go mod tidy
```

2.  运行

```shell
go run .
```

#### Docker部署

以Ubuntu为例

##### 卸载旧版本

```shell
for pkg in docker.io docker-doc docker-compose podman-docker containerd runc; do apt-get remove $pkg; done
```

##### 安装依赖

```shell
apt-get update
apt-get install ca-certificates curl gnupg
```

##### 添加GPG公钥

```shell
install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /etc/apt/keyrings/docker.gpg
sudo chmod a+r /etc/apt/keyrings/docker.gpg
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://mirrors.tuna.tsinghua.edu.cn/docker-ce/linux/ubuntu \
  "$(. /etc/os-release && echo "$VERSION_CODENAME")" stable" | \
  tee /etc/apt/sources.list.d/docker.list > /dev/null
```

##### 安装docker

```shell
apt-get update
apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

##### 编写Dockerfile

> dockerfile参考文档：https://www.jb51.net/server/293884fmf.htm

```dockerfile
# 使用官方Go镜像作为构建环境
FROM golang:latest
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum文件，以确保依赖关系一致性
COPY go.mod go.sum ./
# 获取依赖项
RUN go mod download

# 复制项目源代码
COPY . /app

# 构建项目
RUN go build .

EXPOSE 3000

ENTRYPOINT ["./ginblog"]
```
##### 配置ginblog的config

> host: ginblog-mysql 是为了后面容器互通做准备，对应的是mysql容器的name

```yaml
mysql:
  host: ginblog-mysql
  port: 3306
  user: ginblog
  password: 123456
  dbname: ginblog
```

##### 前端web文件夹下axios请求地址，前端推荐使用Nginx部署

```shell
axios.defaults.baseURL = 'http://localhost:3000/api/v1'

# 改为

axios.defaults.baseURL = 'http://线上服务器ip或域名:3000/api/v1'
```

修改地址后，重新打包
```shell
pnpm run build
```

##### 生成镜像

生成`ginblog docker image`

```shell
docker build -t ginblog .
docker run -d -p 3000:3000 --name ginblog ginblog
```

#### 使用说明

1.  xxxx
2.  xxxx
3.  xxxx

