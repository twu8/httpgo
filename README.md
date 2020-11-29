# `httpgo` A simple http server

## Create from Cobra

```
go get -u github.com/spf13/cobra/cobra
mkdir -p httpgo && cd httpgo
cobra init --pkg-name httpgo
go mod init httpgo

```

## Build

`go build`

## Test

`go test ./api -v`

## Run

`./httpgo`

## APIs

```bash
http://localhost:8080/
http://localhost:8080/api/echo?message=hello
http://localhost:8080/api/books
```




# Docker 

## Build and push

build:

```bash
docker build -t httpgo:1.0.0 .
docker images
docker image rm 6d8772fbd285
docker image rm httpgo:1.0.0
```

push:

```bash
docker tag httpgo:1.0.1 ouyang8/httpgo:1.0.1
```

```bash
docker login -u ouyang8
Password:
Login Succeeded

docker push ouyang8/httpgo:1.0.1
The push refers to repository [docker.io/ouyang8/httpgo]
```
Docker repository

`https://hub.docker.com/repository/docker/ouyang8/httpgo`

## Running Docker image locally

```bash
docker run -it -p 8086:8086 httpgo:1.0.1
docker run -it -e "PORT=9090" -p 9090:9090 httpgo:1.0.1
docker ps --all

docker run --name=httpgo -d -p 8080:8080 httpgo:1.0.1
docker ps
docker stats
docker stop
docker start 
docker kill httpgo
docker rm httpgo
```

## Reduce the Docker image size

```bash
GOOS=linux GOARCH=amd64 go build
docker build -t httpgo:1.0.2-alpine .
docker images
docker tag httpgo:1.0.2-alpine ouyang8/httpgo:1.0.2-alpine
docker push ouyang8/httpgo:1.0.2-alpine
```

## Docker Compose

```bash
$ docker-compose build
$ docker images
$ docker-compose up -d
$ docker-compose kill
$ docker-compose rm
```