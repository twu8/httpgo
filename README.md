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

`./httpgo 8086`

## APIs

```bash
http://localhost:8080/
http://localhost:8080/api/echo?message=hello
http://localhost:8080/api/books
```
## Response

### Delay send response back

For example, following delay 5 seconds, and 350ms
```bash
http://localhost:8080/response/deply/5s
http://localhost:8080/response/deply/350ms
```

### Echo back request headers in response body

```bash
http://localhost:8080/response/header/

curl  -H "abc: def" http://localhost:8086/response/header/
user-agent=curl/7.64.1; accept=*/*; abc=def%

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
docker run -it -p 8088:8088 httpgo:1.0.1
docker run -it -e "PORT=9090" -p 9090:9090 httpgo:1.0.1
docker run -it -e "PORT=9090" -p 9090:9090 httpgo:1.0.4-alpine
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
docker-compose build
docker images
docker-compose up -d
docker-compose kill
docker-compose rm
```

## Building with Makefile

The `Makefile` in the project root allows for convenient local builds and includes version information (version, build time, commit hash) into the compiled binary. This version information is accessible via the `/version` endpoint.

Here are the common `make` commands:

*   **Build the application:**
    ```bash
    make build
    ```
    This compiles the `httpgo` binary. The default version is "dev".

*   **Build with a specific version:**
    You can specify a version by setting the `VERSION` variable:
    ```bash
    make build VERSION=v1.0.0
    ```

*   **Run the application:**
    This command will first build the application (if not already built or if sources have changed) and then run it.
    ```bash
    make run
    ```
    To run with a specific version:
    ```bash
    make run VERSION=v1.0.1
    ```

*   **Clean build artifacts:**
    This removes the compiled `httpgo` binary.
    ```bash
    make clean
    ```

*   **Show version information:**
    This command displays the version, build time, and commit hash that will be compiled into the binary without actually performing a full build.
    ```bash
    make show_version
    ```
    To see what a specific version would show:
    ```bash
    make show_version VERSION=v1.2.3
    ```