# goecho

    go get github.com/daved/goecho/...

Package echo provides a simple udp echo example and dockerfile. An https request
is also sent before the echo server runs in order to ensure that the docker
container has working access to the net lib and ca-certs.

## Usage

### Run

```
goecho
```

### Build (local)

```
go build
```

### Build (container)

```
docker build -t {namespace}/goecho .
```

### Run (container)
```
docker run -i -p {port}:25000 {namespace}/goecho:latest
```
