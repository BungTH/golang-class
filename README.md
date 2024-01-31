# How to Build

## Normal Build

```sh
go build main.go
```

## For gin-demo
### require ldflags for dynamic buildcommit and buildtime var

```sh
go build \
-ldflags "-X main.buildcommit=`git rev-parse --short HEAD` \
-X main.buildtime=`date "+%Y-%m-%dT%H:%M:%S%Z:00"`" \
-o main
```

# How to Run

## Normal Run

```sh
go run main.go
```

## For gin-demo

```sh
./main
```

Please open each sub folder seperately to run the program.
