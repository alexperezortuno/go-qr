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

# Run SMS
```shell
./build/go_qr sms -n +1234567890 -m "Hello World" -o qr_sms.png
```

# Run Geo
```shell
  ./build/go_qr geo -x 37.7749 -y -122.4194 -o qr_geo.png
```

## Run Deeplink
```shell
./build/go_qr deeplink -d test -s 'parameters?test1=1&test2=2' -o qr_deeplink.png
```

## Run Crypto
```shell
./build/go_qr crypto -c BTC -a test123456asdfghj -d test -m 0.001 -o qr_crypto.png
```

## Run Wifi
```shell
./build/go_qr wifi -s test -p test123456 -t WPA -o qr_wifi.png
```

## Run Calendar
```shell
./build/go_qr calendar -e "Test" -a "Test address" -s "2024-01-01 00:00:00" -d "2024-01-01 01:00:00" -o qr_calendar.png
```
