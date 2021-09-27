set GOOS=linux
set GOARCH=amd64
go build -gcflags=-m -ldflags="-w -s" -o tmp/mcss main.go
upx tmp/mcss
