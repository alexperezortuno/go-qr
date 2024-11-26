# go-qr
qr codes with go

## Build
```shell
 go build -a -o build/go_qr main.go
```

## Build windows
```PowerShell
$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o build/go_qr.exe main.go
```
