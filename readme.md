# Try to play golang tags

## How to test

```sh
go test -tags=unit ./...
go test -tags=integration ./...
```

## Send Flag on Build

```sh
go build -ldflags="-X <var>" <file.go>
```
