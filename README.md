
## Overview

### 源码部署安装

1. 安装MySql
2. 安装Redis
3. 安装Golang
4. 安装Node.js
5. 项目根目录执行`npm i`
6. 项目根目录执行`export GOPROXY=https://goproxy.cn go mod download`
7. 项目根目录执行`go build -o bin/enroll main/cmd.go`
8. 项目bin目录修改`config.yaml`
8. 项目bin执行`./enroll`




