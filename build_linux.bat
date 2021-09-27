set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -gcflags=-m -ldflags="-w -s" -o tmp/mcss.linux main.go
upx tmp/mcss.linux
