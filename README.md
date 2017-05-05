# goecho

    go get github.com/daved/goecho

goecho is a simple udp echo example. An https request is also sent before the 
echo server runs in order to ensure that a docker container built from the
provided dockerfile has working access to the net lib and ca-certs.

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

### Interact (telnet)

```
telnet 0 25000
```

### Interact (netcat)

```
nc 0 25000
```
