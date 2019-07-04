# aerospike-oom

## Using the pre-build binary
```
./aerospike-oom -host localhost -port 3000 -namespace test -elements 40000
```

## Using Docker

```
docker build -t aerospike-oom .
docker run --network host -ti aerospike-oom:latest -host localhost -port 3000 -namespace test -elements 40000
```

## Using local Go environment (tested with Go 1.12)

Run this command

```
go run main.go -host localhost -port 3000 -namespace test -elements 40000
```

Then wait a few seconds and the OOM killer kills the `asd` process.

