

run:
	go run ./cmd/*go

up:
	docker-compose up -d

down:
	docker-compose down

net:
	@docker inspect mysql8 | grep IPAddress

mysql:
	@docker exec -it mysql8 bash
# mysql -uroot -p

redis:
	@docker exec -it redis bash
# redis-cli


# === 部署 k8s
build:

# 步骤 1
# 把上次编译的东西删掉
# @rm webook || true

# 步骤 2
# 运行一下 go mod tidy，防止 go.sum 文件不对，编译失败
# @go mod tidy

# 步骤 3
# 指定编译成在 AMD64 架构的 linux 操作系统上运行的可执行文件，名字叫做 webook
	@GOOS=linux GOARCH=amd64 go build -tags=k8s -o webook ./cmd/*go

# 步骤 4
# 把上次打包的镜像删掉
# @docker rmi -f shgqmrf/webook:v0.0.1

# 步骤 5
# 这里可以随便改这个标签，记得对应的 k8s 部署里面也要改
	@docker build -t shgqmrf/webook:v0.0.1 .

# 步骤 6
# 登录 docker
# @docker login

# 步骤 7
# 推送镜像到 docker hub
# @docker push shgqmrf/webook:v0.0.1 


.PHONY: run up down net build
