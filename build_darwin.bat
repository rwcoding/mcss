set CGO_ENABLED=0
set GOOS=darwin
set GOARCH=amd64
go build -gcflags=-m -ldflags="-w -s" -o tmp/mcss.darwin main.go
upx tmp/mcss.darwin
