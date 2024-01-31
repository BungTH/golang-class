# How to Build

```sh
go build \
-ldflags "-X main.buildcommit=`git rev-parse --short HEAD` \
-X main.buildtime=`date "+%Y-%m-%dT%H:%M:%S%Z:00"`" \
-o main
```

# How to Run

```sh
./main
```
